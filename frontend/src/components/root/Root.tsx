import React, { useState, useCallback } from 'react';
import { FetchBaseQueryError } from '@reduxjs/toolkit/dist/query';
import block from 'bem-cn-lite';
import { useNavigate } from 'react-router-dom';

import { AsideNavigation } from '../AsideNavigation/AsideNavigation';
import { Page } from '../../pages';

import { TopBar } from '../TopBar/TopBar';

// import { Spinner } from '../spinner';

import './Root.scss';

export const cn = block('Root');


export const Root: React.FC = () => {
    const navigate = useNavigate();

    const navigateCallback = useCallback((path: string) => {
        navigate(path);
    }, [navigate]);

    // const { data: env, isLoading, error } = useGetEnvironmentQuery();


    return (
        <div className={cn()}>
            <AsideNavigation navigateCallback={navigateCallback} />
            <div className={cn('MainInfo')}>
                <TopBar />
                <Page />
            </div>
        </div>
    );
};
