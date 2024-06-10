import React from 'react';

import { cn } from '@bem-react/classname';
import {
    Routes,
    Route,
    Navigate,
} from 'react-router-dom';

import { HomePage } from './home';
import { MintsPage } from './mints';
// import { useEnvironment } from '../components';

import './Page.css';
import 'font-awesome/css/font-awesome.css';

export const cnPage = cn('Page');

export const routes = {
    home: '/',
    mints: '/mints',
    profile: '/profile',
};

export const Page: React.FC = () => {
    // const env = useEnvironment();
    const userHasPermissions = true; // get from env
    if (!userHasPermissions) {
        return (
            <div style={{ textAlign: 'center' }}>
                Access denied.
            </div>
        );
    }

    return (
        <>
            <div className={cnPage()} id="main-page">
                <Routes>
                    <Route path={routes.home} element={<HomePage />} />
                    <Route path={routes.mints} element={<MintsPage />} />
                    <Route path="*" element={<Navigate to={routes.home} />} /> { /* TODO 404 */}
                </Routes>
            </div>
        </>
    );
};
