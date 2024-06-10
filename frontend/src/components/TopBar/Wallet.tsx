import React, { useState, useMemo, useRef, useEffect, useCallback } from 'react';

import { useNavigate, useSearchParams } from "react-router-dom";
import { jwtDecode } from 'jwt-decode';
import { useCookies } from 'react-cookie';
import block from 'bem-cn-lite';
import base58 from 'bs58';

import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { Button } from '@mui/joy';
import { useWallet } from '@solana/wallet-adapter-react';
import { useWalletMultiButton } from '@solana/wallet-adapter-base-ui';
import { useWalletModal } from '@solana/wallet-adapter-react-ui';

import { useUserInitMutation, useUserVerifyMutation } from '../../features/user/userApi';
import { useErrorToaster } from '../Toastr';

import './TopBar.scss';

const cn = block('TopBar');

const labels = {
    'change-wallet': 'Change wallet',
    connecting: 'Connecting ...',
    'copy-address': 'Copy address',
    copied: 'Copied',
    disconnect: 'Disconnect',
    'has-wallet': 'Connecting ...',
    'no-wallet': 'Connect Wallet',
} as const;

export const WalletComponent: React.FC = () => {
    const navigate = useNavigate();
    const [searchParams, setSearchParams] = useSearchParams();

    const { setVisible: setModalVisible } = useWalletModal();

    const [cookies, setCookie, removeCookie] = useCookies(['jwt'])

    const [menuOpen, setMenuOpen] = useState(false);
    const [copied, setCopied] = useState(false);
    const [approvedWallet, setApprovedWallet] = useState(false);
    const [disconnectingWallet, setDisconnectingWallet] = useState(false);
    const [siginingWallet, setSigningWallet] = useState(false);
    const [nonce, setNonce] = useState<string | undefined>();

    const { connect, disconnect, signMessage } = useWallet();
    const [useUserInitQuery, { data: userInitData, isLoading: isInitLoading, isSuccess, error: initError }] = useUserInitMutation();
    const [useUserVerifyQuery, { data: userVerifyData, isLoading: isVerifying, error: verifyError }] = useUserVerifyMutation();

    useErrorToaster(initError);

    useEffect(() => {
        if (verifyError) {
            localStorage.setItem('hasNoAccess', 'yes');
            navigate("/");
            // setMenuOpen(false);
            // setApprovedWallet(false);
            // setDisconnectingWallet(true);
            // try { disconnect() } catch(e) {}
        }
    }, [verifyError])

    // useErrorToaster(verifyError, verifyErrorHeader, verifyErrorText);

    const { buttonState, publicKey } = useWalletMultiButton({
        onSelectWallet() {
            setModalVisible(true);
        },
    });

    useEffect(() => {
        if (cookies.jwt && publicKey && buttonState === 'connected') {
            const decoded = jwtDecode(cookies.jwt) as Record<string, string | number>;
            setApprovedWallet(decoded.id === publicKey.toBase58());
        }
    }, [cookies, publicKey, buttonState]);

    const ref = useRef<HTMLUListElement>(null);
    useEffect(() => {
        const listener = (event: MouseEvent | TouchEvent) => {
            const node = ref.current;

            if (!node || node.contains(event.target as Node)) return;

            setMenuOpen(false);
        };

        document.addEventListener('mousedown', listener);
        document.addEventListener('touchstart', listener);

        return () => {
            document.removeEventListener('mousedown', listener);
            document.removeEventListener('touchstart', listener);
        };
    }, []);

    useEffect(() => {
        if (nonce && publicKey && !siginingWallet && !approvedWallet && signMessage) {
            const message = new TextEncoder().encode(nonce);
            setSigningWallet(true);
            signMessage(message).then(signature => {
                useUserVerifyQuery({
                    address: publicKey.toBase58(),
                    signature: base58.encode(signature),
                    ref_code: searchParams.get('refferalCode') || undefined
                });
            }).catch(() => {
                setSigningWallet(false);
            });
        }
    }, [nonce, publicKey, siginingWallet, signMessage, approvedWallet]);

    useEffect(() => {
        if (isSuccess && userInitData) {
            setNonce(userInitData.nonce);
        }
    }, [isSuccess, userInitData]);

    useEffect(() => {
        if (!isVerifying && userVerifyData) {
            setCookie('jwt', userVerifyData.token, { path: '/', expires: new Date(userVerifyData.expire) });
            localStorage.setItem('token', userVerifyData.token);
            navigate("/");
            setSigningWallet(false);
            setApprovedWallet(true);
        }
    }, [isVerifying, userVerifyData]);

    useEffect(() => {
        if (buttonState === 'connected' && publicKey && !approvedWallet && !disconnectingWallet) {
            if (!cookies?.jwt || publicKey.toBase58() !== (jwtDecode(cookies.jwt) as Record<string, string | number>).id) {
                setSigningWallet(false);
                useUserInitQuery(publicKey.toBase58());
            }
        } else {
            setNonce(undefined);
        }
        if (buttonState === 'connecting') {
            setApprovedWallet(false);
            setSigningWallet(false);
        }

        if (buttonState === 'has-wallet' && !disconnectingWallet) {
            connect();
        }

        if (buttonState === 'disconnecting') {
            removeCookie('jwt');
            localStorage.removeItem('token');
            localStorage.removeItem('hasNoAccess');
            navigate("/");
        }
        if (buttonState === 'no-wallet') {
            removeCookie('jwt');
            localStorage.removeItem('token');
            localStorage.removeItem('hasNoAccess');
            // navigate("/");
            setDisconnectingWallet(false);
        }
    }, [buttonState, connect, disconnectingWallet, setDisconnectingWallet, approvedWallet, publicKey, setSigningWallet, cookies]);

    const content = useMemo(() => {
        if (buttonState === 'connected' && publicKey && approvedWallet) {
            const base58 = publicKey.toBase58();
            return <>
                <span>{base58.slice(0, 4) + '..' + base58.slice(-4)}</span>
                <AccountCircleIcon />
            </>
        } else if (buttonState === 'connected' && !approvedWallet || buttonState === 'connecting' || buttonState === 'has-wallet') {
            return labels['connecting'];
        } else {
            return labels['no-wallet'];
        }
    }, [buttonState, labels, publicKey, approvedWallet]);

    const onDropdownCopy = useCallback(async () => {
        await navigator.clipboard.writeText(publicKey?.toBase58() || '');
        setCopied(true);
        setTimeout(() => setCopied(false), 400);
    }, []);

    const onDropdownChangeWallet = useCallback(() => {
        setModalVisible(true);
        setMenuOpen(false);
    }, []);

    const onDropdownDisconnectWallet = useCallback(() => {
        setDisconnectingWallet(true);
        try { disconnect() } catch(e) {}
        removeCookie('jwt');
        localStorage.removeItem('token');
        localStorage.removeItem('hasNoAccess');
        navigate("/");
        setMenuOpen(false);
        setApprovedWallet(false);
    }, []);

    const onConnectButton = useCallback(() => {
        if (buttonState === 'no-wallet') {
            setModalVisible(true)
        } else {
            setMenuOpen(true)
        }
    }, [buttonState]);

    return (
        <div className="wallet-adapter-dropdown">
            <Button
                className={cn('ConnectButton')}
                variant="plain"
                onClick={onConnectButton}
            >
                {content}
            </Button>
            <ul
                aria-label="dropdown-list"
                className={`wallet-adapter-dropdown-list ${menuOpen && 'wallet-adapter-dropdown-list-active'}`}
                ref={ref}
                role="menu"
            >
                {publicKey ? (
                    <li
                        className="wallet-adapter-dropdown-list-item"
                        onClick={onDropdownCopy}
                        role="menuitem"
                    >
                        {copied ? labels['copied'] : labels['copy-address']}
                    </li>
                ) : null}
                <li
                    className="wallet-adapter-dropdown-list-item"
                    onClick={onDropdownChangeWallet}
                    role="menuitem"
                >
                    {labels['change-wallet']}
                </li>
                <li
                    className="wallet-adapter-dropdown-list-item"
                    onClick={onDropdownDisconnectWallet}
                    role="menuitem"
                >
                    {labels['disconnect']}
                </li>
            </ul>
        </div>
    )
}
