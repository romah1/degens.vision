import React from "react";

import block from "bem-cn-lite";

import "./styles.scss";

const cn = block("PlayPauseIcon");

type IconProps = {
    onClick?: () => void;
};

export const PauseIcon: React.FC<IconProps> = ({ onClick }) => {
    return (
        <svg
            width="30"
            height="30"
            viewBox="0 0 30 30"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            onClick={onClick}
            className={cn()}
        >
            <g clipPath="url(#clip0_472_1699)">
                <path
                    d="M15 26.25C21.2132 26.25 26.25 21.2132 26.25 15C26.25 8.7868 21.2132 3.75 15 3.75C8.7868 3.75 3.75 8.7868 3.75 15C3.75 21.2132 8.7868 26.25 15 26.25Z"
                    stroke="#0ABAB5"
                    strokeWidth="1.5"
                    strokeMiterlimit="10"
                />
                <path
                    d="M12.1875 11.25V18.75"
                    stroke="#0ABAB5"
                    strokeWidth="1.5"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                />
                <path
                    d="M17.8125 11.25V18.75"
                    stroke="#0ABAB5"
                    strokeWidth="1.5"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                />
            </g>
            <defs>
                <clipPath id="clip0_472_1699">
                    <rect width="30" height="30" fill="white" />
                </clipPath>
            </defs>
        </svg>
    );
};

export const PlayIcon: React.FC<IconProps> = ({ onClick }) => {
    return (
        <svg
            width="30"
            height="30"
            viewBox="0 0 30 30"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            onClick={onClick}
            className={cn()}
        >
            <g clipPath="url(#clip0_924_24)">
                <path
                    d="M15 26.25C21.2132 26.25 26.25 21.2132 26.25 15C26.25 8.7868 21.2132 3.75 15 3.75C8.7868 3.75 3.75 8.7868 3.75 15C3.75 21.2132 8.7868 26.25 15 26.25Z"
                    stroke="#0ABAB5"
                    strokeWidth="1.5"
                    strokeMiterlimit="10"
                />
                <path
                    d="M12.1875 18.75V11.25L18.8 15L12.1875 18.75Z"
                    stroke="#0ABAB5"
                    strokeWidth="1.5"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                />
            </g>
            <defs>
                <clipPath id="clip0_924_24">
                    <rect width="30" height="30" fill="white" />
                </clipPath>
            </defs>
        </svg>
    );
};
