import React, { useCallback, useMemo, memo } from 'react';

import {
    Button,
} from '@mui/joy';
import block from 'bem-cn-lite';

import { CrownIcon, DiscordIcon, FireIcon, LogoIcon, LogoTextIcon, MenuClosedIcon, MenuOpenedIcon, MintsIcon, StatisticIcon, StonksIcon, TargetIcon, TwitterIcon, WhitepaperIcon } from '../icons';


import './AsideNavigation.scss';

const cn = block('AsideNavigation');

type ActionButtonProps = {
    name: string;
    icon: React.ReactNode;
    onClick: () => void;
    disabled?: boolean;
}

const ActionButton: React.FC<ActionButtonProps> = (props) => {
    const { name, icon, onClick, disabled } = props;

    return (
        <Button
            disabled={disabled}
            className={cn('ActionButton')}
            variant="plain"
            startDecorator={icon}
            onClick={onClick}
        >
            {name}
        </Button>
    )
}

type Props = {
    navigateCallback: (url: string) => void;
}

const Sidebar: React.FC<Props> = ({ navigateCallback }) => {

    const navigateCallbackNew = useCallback((url: string) => () => {
        navigateCallback(url);
    }, [navigateCallback]);

    const actionButtons = useMemo(() => {
        return [
            {
                name: 'Mints',
                icon: <MintsIcon />,
                onClick: navigateCallbackNew('/mints'),
                disabled: false
            },
            // {
            //     name: 'Statistic',
            //     icon: <StatisticIcon />,
            //     onClick: navigateCallbackNew('/statistic'),
            //     disabled: true
            // },
            // {
            //     name: 'Fire',
            //     icon: <FireIcon />,
            //     onClick: navigateCallbackNew('/fire'),
            //     disabled: true
            // },
            // {
            //     name: 'Target',
            //     icon: <TargetIcon />,
            //     onClick: navigateCallbackNew('/target'),
            //     disabled: true
            // },
            // {
            //     name: 'Crown',
            //     icon: <CrownIcon />,
            //     onClick: navigateCallbackNew('/crown'),
            //     disabled: true
            // },
            // {
            //     name: 'Stonks',
            //     icon: <StonksIcon />,
            //     onClick: navigateCallbackNew('/stonks'),
            //     disabled: true
            // },
        ]
    }, [navigateCallbackNew]);

    return (
        <>
            <div className={cn('Logo')} onClick={navigateCallbackNew('/')}>
                <LogoIcon />
                <LogoTextIcon />
            </div>
            <div className={cn('OtherButtons')}>
                <div className={cn('Actions')}>
                    {actionButtons.map((button, index) => (
                        <ActionButton key={index} {...button} />
                    ))}
                </div>
                <div className={cn('Socials')}>
                    <Button
                        className={cn('ActionButton')}
                        variant="plain"
                        component="a"
                        target="_blank"
                        rel="noopener noreferrer"
                        title="Our Twitter"
                        href="https://twitter.com/DegensVision"
                        startDecorator={<TwitterIcon />}
                    >
                        Twitter
                    </Button>
                    <Button
                        className={cn('ActionButton')}
                        variant="plain"
                        component="a"
                        target="_blank"
                        rel="noopener noreferrer"
                        title="Our Discord"
                        href="https://discord.gg/degensvision"
                        startDecorator={<DiscordIcon />}
                    >
                        Discord
                    </Button>
                    <Button
                        className={cn('ActionButton')}
                        variant="plain"
                        component="a"
                        target="_blank"
                        rel="noopener noreferrer"
                        title="Our Whitepaper"
                        href="https://whitepaper.degens.vision/"
                        startDecorator={<WhitepaperIcon />}
                    >
                        Whitepaper
                    </Button>
                </div>
            </div>
        </>
    )
}


export const AsideNavigation: React.FC<Props> = memo(({ navigateCallback }) => {
    return (
        <div className={cn('Sidebar')}>
            <Sidebar navigateCallback={navigateCallback} />
        </div>
    );
})
