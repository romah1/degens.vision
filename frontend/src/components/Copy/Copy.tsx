import React, { useState, useEffect, useCallback } from "react";

import { CopyIcon, CopiedIcon } from '../icons';

import block from 'bem-cn-lite';

import './Copy.scss';

const cn = block('Copy');

type CopyComponentProps = {
    className?: string;
    value?: string;
    copyIconClassName?: string;
    copyValue?: string;
}

export const CopyComponent: React.FC<CopyComponentProps> = (props) => {
    const { className, value, copyIconClassName, copyValue } = props;
    const [copySuccess, setCopySuccess] = useState(false);


    useEffect(() => {
        const timeout = setTimeout(() => setCopySuccess(false), 1000);
        return () => clearTimeout(timeout);
    }, [copySuccess]);

    const handleCopy = useCallback(() => {
        if (copyValue) {
            navigator.clipboard.writeText(copyValue);
            setCopySuccess(true);
        }
    }, [copyValue]);

    return (
        <div className={cn() + ' ' + (className || '')}>
            <a className={cn('value')} href={`https://solscan.io/account/${copyValue}`} target="_blank" rel="noopener noreferrer">{value}</a>
            <div onClick={handleCopy} className={cn('copy')}>
                {!copySuccess ? <CopyIcon /> : <CopiedIcon /> }
            </div>
        </div>
    )
}
