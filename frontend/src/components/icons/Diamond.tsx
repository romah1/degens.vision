import React from "react";

export const DiamondIcon: React.FC = () => {
    return (
        <svg
            width="20"
            height="20"
            viewBox="0 0 20 20"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
        >
            <g clipPath="url(#clip0_472_1823)">
                <path
                    d="M5.625 3.125H14.375L18.75 8.125L10 17.5L1.25 8.125L5.625 3.125Z"
                    stroke="white"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                />
                <path
                    d="M13.75 8.125L10 17.5L6.25 8.125L10 3.125L13.75 8.125Z"
                    stroke="white"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                />
                <path
                    d="M1.25 8.125H18.75"
                    stroke="white"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                />
            </g>
            <defs>
                <clipPath id="clip0_472_1823">
                    <rect width="20" height="20" fill="white" />
                </clipPath>
            </defs>
        </svg>
    );
};
