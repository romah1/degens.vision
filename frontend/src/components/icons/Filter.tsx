import React from "react";

import block from 'bem-cn-lite';

import './styles.scss';

const cn = block('FilterIcon');

type IconProps = {
    onClick?: () => void;
}

export const FilterIcon: React.FC<IconProps> = ({ onClick }) => {
    const [isActive, setIsActive] = React.useState(false);

    const handleClick = React.useCallback(() => {
        setIsActive(!isActive);
        if (onClick) onClick();
    }, [isActive]);

    return (
        <svg
            width="30"
            height="30"
            viewBox="0 0 30 30"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            className={cn() + (isActive ? ' Active' : '')}
            onClick={handleClick}
        >
            <g clipPath="url(#clip0_472_544)">
                <path
                    d="M3.99757 7.19297C3.87558 7.05872 3.79516 6.89196 3.76605 6.71291C3.73695 6.53386 3.76041 6.35022 3.8336 6.18424C3.90679 6.01826 4.02656 5.87709 4.17839 5.77782C4.33021 5.67856 4.50758 5.62547 4.68897 5.625H25.314C25.4956 5.62502 25.6733 5.67778 25.8254 5.77686C25.9776 5.87595 26.0977 6.01711 26.1712 6.18317C26.2447 6.34924 26.2684 6.53307 26.2393 6.71233C26.2103 6.8916 26.1298 7.05857 26.0077 7.19297L18.0694 15.6668C17.9067 15.8406 17.8162 16.0698 17.8163 16.3078V22.8105C17.8164 22.9649 17.7784 23.117 17.7056 23.2531C17.6328 23.3893 17.5276 23.5053 17.3991 23.591L13.6491 26.0906C13.5081 26.1854 13.3439 26.24 13.1741 26.2486C13.0044 26.2573 12.8355 26.2196 12.6856 26.1397C12.5356 26.0598 12.4102 25.9406 12.3227 25.7949C12.2352 25.6492 12.189 25.4824 12.189 25.3125V16.3078C12.189 16.0698 12.0986 15.8406 11.9358 15.6668L3.99757 7.19297Z"
                    stroke="white"
                    strokeWidth="1.5"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                />
            </g>
            <defs>
                <clipPath id="clip0_472_544">
                    <rect width="30" height="30" fill="white" />
                </clipPath>
            </defs>
        </svg>
    );
};
