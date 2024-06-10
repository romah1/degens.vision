import { UIRuntimeEnvName } from './interfaces/common';
import * as buffer from 'buffer/index';

if (typeof (window as any).global === "undefined"){
    (window as any).global = window;
}
if (typeof (window as any).Buffer === "undefined") {
    (window as any).Buffer = buffer.Buffer;
}

// Change this
type Environment = {
    name: string;
    apiUrl: string;
    uiRelease?: string;
}

interface WindowWithEnv {
    dvEnvironment: Environment;
}

export const dvEnv = (window as unknown as WindowWithEnv).dvEnvironment;

export function getEnvName(): UIRuntimeEnvName {
    const name = dvEnv?.name || '';
    if (name === 'prod') {
        return UIRuntimeEnvName.PROD;
    }
    if (name === 'test') {
        return UIRuntimeEnvName.TEST;
    }
    return name as UIRuntimeEnvName;
}

export const isProd = () => dvEnv?.name === 'prod';
export const isTest = () => dvEnv?.name === 'test';

export const baseWebsocketUrl = 'wss://api.degens.vision';
export const baseApiUrl = 'https://api.degens.vision';

export const defaultFetchParams: RequestInit = {
    credentials: 'include',
    mode: 'cors',
    headers: { 'Content-Type': 'application/json' },
};
