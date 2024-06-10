import React from "react";

import block from 'bem-cn-lite';

import './Container.scss';

const cn = block('Container');

interface ContainerProps {
    children: React.ReactNode;
}

export const Container: React.FC<ContainerProps> = ({children}) => {
    return (
        <div className={cn()}>
            {children}
        </div>
    )
}
