import React, { memo } from "react";

import block from 'bem-cn-lite';

import './CollectionContainer.scss';

const cn = block('CollectionContainer');

interface CollectionContainerProps {
    children: React.ReactNode;
}

export const CollectionContainer: React.FC<CollectionContainerProps> = memo(({children}) => {
    return (
        <div className={cn()}>
            {children}
        </div>
    )
})
