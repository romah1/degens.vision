import React, { useMemo } from 'react';

import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { clusterApiUrl } from '@solana/web3.js';
import { SkeletonTheme } from 'react-loading-skeleton';

import { WalletAdapterNetwork } from '@solana/wallet-adapter-base';
import {
    LedgerWalletAdapter,
    PhantomWalletAdapter,
    SolflareWalletAdapter,
    TorusWalletAdapter,
} from '@solana/wallet-adapter-wallets';
import { ConnectionProvider, WalletProvider } from '@solana/wallet-adapter-react';
import { WalletModalProvider } from '@solana/wallet-adapter-react-ui';

import { BrowserRouter } from 'react-router-dom';

import { Root } from './components/root/Root';
import { store } from './store';

import './index.css';

import { ToastrWrapper } from './components';

import '@solana/wallet-adapter-react-ui/styles.css';
import 'react-loading-skeleton/dist/skeleton.css'

const rootElement = document.getElementById('root');

type State = {
    hasError: boolean;
    stack?: string;
}

class ErrorBoundary extends React.Component<{children?: React.ReactNode}, State> {
    constructor(props: {}) {
        super(props);
        this.state = { hasError: false };
    }
    static getDerivedStateFromError(_error: Error) {
        // Update state so the next render will show the fallback UI.
        return { hasError: true };
    }
    render() {
        if (this.state.hasError) {
            // You can render any custom fallback UI
            return (
                <>
                    <h1>Something went wrong.</h1>
                    <pre>{this.state.stack}</pre>
                </>);
        }
        return this.props.children;
    }
}

const Page: React.FC = () => {
    const solNetwork = WalletAdapterNetwork.Mainnet;
    const endpoint = useMemo(() => clusterApiUrl(solNetwork), [solNetwork]);

    const wallets = useMemo(
        () => [
            new PhantomWalletAdapter(),
            new SolflareWalletAdapter({ network: solNetwork }),
            new TorusWalletAdapter(),
            new LedgerWalletAdapter(),
        ],
        [solNetwork]
    );

    return (
        <ErrorBoundary>
            <Provider store={store}>
                <ConnectionProvider endpoint={'https://solana.magiceden.dev/'}>
                    <WalletProvider wallets={wallets} autoConnect={false}>
                        <WalletModalProvider>
                            <SkeletonTheme baseColor="#202020" highlightColor="#444">
                                <BrowserRouter>
                                    <Root />
                                </BrowserRouter>
                            </SkeletonTheme>
                            <ToastrWrapper position="bottom-right" duration={5000} />
                        </WalletModalProvider>
                    </WalletProvider>
                </ConnectionProvider>
            </Provider>
        </ErrorBoundary>
    );
}

render(
    <Page />,
    rootElement,
);
