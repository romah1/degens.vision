import { LiveMintsItem } from "src/interfaces/collections";
import { baseWebsocketUrl } from "src/config";

type fnDataType = (data: LiveMintsItem) => void;
type fnBoolType = (data: boolean) => void;

export function reconnectingSocket(cl_id: number) {
    let client: WebSocket;
    let isConnected = false;
    let reconnectOnClose = true;
    let messageListeners: fnDataType[] = [];
    let stateChangeListeners: fnBoolType[] = [];

    let currentClient: Record<number, WebSocket> = {};
    let closed = false;

    function on(fn: fnDataType) {
        messageListeners.push(fn);
    }

    function off(fn: fnDataType) {
        messageListeners = messageListeners.filter(l => l !== fn);
    }

    function close() {
        closed = true;
        reconnectOnClose = false
        const cl = currentClient[cl_id];
        if (cl) {
            cl.close();
            delete currentClient[cl_id];
        }
    }

    function onStateChange(fn: fnBoolType) {
        stateChangeListeners.push(fn);
        return () => {
            stateChangeListeners = stateChangeListeners.filter(l => l !== fn);
        };
    }

    function start() {
        const url = `${baseWebsocketUrl}/ws/v1/live_feed?token=${localStorage.getItem('token')}`,
        client = new WebSocket(url);
        currentClient[cl_id] = client;

        client.onopen = () => {
            isConnected = true;
            stateChangeListeners.forEach(fn => fn(true));
        }

        client.onmessage = (event) => {
            messageListeners.forEach(fn => fn(JSON.parse(event.data) as LiveMintsItem));
        }

        client.onerror = (e) => console.error(e);

        client.onclose = () => {
            isConnected = false;
            stateChangeListeners.forEach(fn => fn(false));

            if (!reconnectOnClose) {
                return;
            }

            setTimeout(start, 100);
        }
    }

    start();

    return {
        on,
        off,
        onStateChange,
        close,
        getClient: () => client,
        isConnected: () => isConnected,
    };
}
