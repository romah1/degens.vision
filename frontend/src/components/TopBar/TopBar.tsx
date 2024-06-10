import React, { memo } from 'react';

import block from 'bem-cn-lite';

import { Input } from '@mui/joy';

import { SearchIcon, SolanaIcon } from '../icons';
import { WalletComponent } from './Wallet';

import './TopBar.scss';

const cn = block('TopBar');

export const TopBar: React.FC = memo(() => {
    return (
        <div className={cn()}>
            <div className={cn('search')}>
                <div className={cn('currency')}>
                    <SolanaIcon />
                </div>
                <Input className={cn('search-bar')} type="text" placeholder="Search" startDecorator={<SearchIcon />} />
            </div>
            <div className={cn('user')}>
                <WalletComponent />
            </div>
        </div>
    );
})
