import React from "react";

import block from 'bem-cn-lite';

import {
    Navigate,
} from 'react-router-dom';

import './Home.scss';

const cn = block('HomePage');


export const HomePage: React.FC = () => {
    const token = localStorage.getItem('token');
    const hasNoAccess = localStorage.getItem('hasNoAccess');

    if (token) {
        return <Navigate to="/mints" />;
    }

    return (
        <div className={cn()}>
            {!hasNoAccess && (
                <p className={cn('text')}>
                    Degens.vision now in beta test<br />
                    to access app connect your wallet
                </p>
            )}
            {hasNoAccess && (
                <p className={cn('text')}>
                    You are not degen enough yet..<br />
                    We give invite codes in our <a href="https://twitter.com/DegensVision" target="_blank" rel="noopener noreferrer" className={cn('link')}>Twitter</a><br />
                    Follow to get it
                </p>
            )}

            <img src="https://media.giphy.com/media/3orif0BNmfxTqjtn9K/giphy.gif?cid=ecf05e47b3q1lc156lwk8upxib2smkwu690j6vg992o0traw&ep=v1_gifs_search&rid=giphy.gif&ct=g"  width="480px" height="368px" />
        </div>
    )
}
