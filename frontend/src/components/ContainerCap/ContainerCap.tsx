import React, { memo } from "react";

import block from 'bem-cn-lite';

import './ContainerCap.scss';

const cn = block('ContainerCap');

interface ContainerCapProps {
    name: string;
    nameActionButton?: React.ReactNode;
    actionButton?: React.ReactNode;
    children?: React.ReactNode;
}

export const ContainerCap: React.FC<ContainerCapProps> = memo((props) => {
    const { name, nameActionButton, actionButton, children } = props;

    return (
        <div className={cn()}>
            <div className={cn('Header')}>
                <div className={cn('HeaderText')}>
                    {name}
                    {Boolean(nameActionButton) && nameActionButton}
                </div>
                {actionButton}
            </div>
            {children}
        </div>
    )
})
