import React from "react";

import block from 'bem-cn-lite';
import { Flipper } from "react-flip-toolkit";

import './ContainerBody.scss';

const cn = block('ContainerBody');

interface ContainerBodyProps {
    children: React.ReactNode;
    flipKey: string;
}

export const ContainerBody: React.FC<ContainerBodyProps> = ({ children, flipKey }) => {
    return (
        <Flipper className={cn()} flipKey={flipKey}>
            <ul className={cn('List')}>
                {children}
            </ul>
        </Flipper>
    )
}
