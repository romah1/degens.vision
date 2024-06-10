import React, { useState, useCallback, useEffect, useMemo } from "react";

import block from 'bem-cn-lite';
import Skeleton from 'react-loading-skeleton';
import CountUp from 'react-countup';
import moment from "moment";

import { Flipped } from "react-flip-toolkit";

import './NFTItem.scss';

const cn = block('NFTItem');

interface NFTItemProps {
    name: string;
    img: string;
    price: number;
    platformIcon: React.ReactNode;
    mintsAmount?: number;
    mintedAmount?: number;
    supply?: number;
    isLoading?: boolean;
    onClick?: () => void;
    collectionKey: string;
    onIconClick?: (e: React.MouseEvent<HTMLElement>) => void;

    percents?: number;
    startsMintsAmount: number;
    startsMintedAmount: number;
    startsPercents: number;
    date: string;
}

interface TimeCounterProps {
    date: string;
}

const TimeCounter: React.FC<TimeCounterProps> = props => {
    const [count, setCount] = useState(Math.floor(moment.duration(moment(new Date()).diff(moment(props.date))).asSeconds()));

    useEffect(() => {
        const interval = setInterval(() => {
            setCount(count + 1);
        }, 1000);

        return () => clearInterval(interval);
    }, [count]);

    const buyTimeAgo = useMemo(() => {
        if (count > 60 && count < 120) {
            return `a minute ago`
        } else if (count > 60) {
            return `${Math.floor(count / 60)} minutes ago`
        } else if (count > 1) {
            return `${count} seconds ago`
        } else {
            return `just now`
        }
    }, [count])

    return (
        <div className={cn('buyTimeInner')}>
            <div className={cn('buyTimeHeader')}>MINT</div>
            {buyTimeAgo}
        </div>
    )
}

export const NFTItem: React.FC<NFTItemProps> = props => {
    const {
        collectionKey,
        name,
        img,
        price,
        platformIcon,
        mintsAmount,
        mintedAmount,
        supply,
        isLoading: isLoadingNft,
        onClick,
        percents : nftPercents,
        startsMintsAmount,
        startsMintedAmount,
        startsPercents,
        onIconClick,
        date,
    } = props;

    const isCollection = mintsAmount !== undefined && mintedAmount !== undefined && supply !== undefined;
    const percents = nftPercents || 0;
    const pr = price.toFixed(3) === '0.000' ? '< 0.001' : parseFloat(price.toFixed(3));

    const [loadingBarWidth, setLoadingBarWidth] = React.useState(0);
    const [isLoadingImage, setLoadingImage] = React.useState(true);
    const [isCounting, setIsCounting] = React.useState(false);

    const isLoading = isLoadingNft || isLoadingImage;

    const handleOnLoad = useCallback(() => {
        setLoadingImage(false);
    }, []);

    const handleOnLoadStart = useCallback(() => {
        const timeout = setTimeout(() => setLoadingImage(false), 0);
        return () => clearTimeout(timeout);
    }, []);

    const onExit = React.useCallback((element: HTMLElement, index: number, removeElement: () => void) => {
        setIsCounting(false);
        setLoadingBarWidth(0);
        removeElement();
    }, []);

    const onAppear = React.useCallback((element: HTMLElement) => {
        element.style.opacity = '0';
        setLoadingBarWidth(startsPercents)
        setTimeout(() => {
            element.style.transition = 'opacity 0.3s';
            element.style.opacity = '1';
            setLoadingBarWidth(percents)
        }, 0);
    }, [startsPercents, percents]);

    // const onStart = React.useCallback((element: HTMLElement) => {
    //     setHasOpacity(true);
    // }, []);

    useEffect(() => {
        setLoadingBarWidth(percents);
        if (startsMintsAmount !== mintsAmount && mintsAmount !== undefined) {
            setIsCounting(true);
        }
    }, [startsMintsAmount, mintsAmount, percents]);

    const finishCounting = React.useCallback(() => {
        setIsCounting(false);
    }, []);

    return (
        <Flipped
            key={collectionKey}
            onExit={onExit}
            onAppear={onAppear}
            flipId={collectionKey}
        >
            <li className={cn() + (isCounting ? ' Animating' : '') + (isLoadingNft ? ' Loading' : '')} onClick={onClick}>
                <div className={cn('Inner')}>
                    <div className={cn('Image') + (!isCollection ? ' HoverImage' : '')} onClick={onIconClick}>
                        {isLoading && <Skeleton className={cn('Image')} />}
                        {!isLoadingNft &&
                            <img
                                onClick={onIconClick}
                                src={img}
                                style={{ opacity: isLoading ? 0 : 1 }}
                                className={cn('Image')}
                                onLoad={handleOnLoad}
                                onLoadStart={handleOnLoadStart}
                            />
                        }
                    </div>

                    <div className={cn('BaseInfo')}>
                        <div className={cn('LeftInfo') + (!isCollection ? ' Wide' : '')}>
                            <div className={cn('Name')}>
                                {isLoadingNft ?
                                    <Skeleton /> :
                                    name
                                }
                            </div>
                            {isCollection &&
                                <div className={cn('CollectionInfo')}>

                                    <div className={cn('ItemText')}>
                                        {isLoadingNft ?
                                            <Skeleton count={2} /> :
                                            <>
                                                <div>
                                                    <CountUp
                                                        onEnd={finishCounting}
                                                        start={startsMintsAmount}
                                                        end={mintsAmount}
                                                        duration={1}
                                                        suffix=" mints"
                                                    />
                                                </div>
                                                <div>{price === 0 ? 'Free Mint' : `${pr} sol`}</div>
                                            </>
                                        }
                                    </div>
                                    <div className={cn('PlatformIcon')}>
                                        {isLoadingNft ?
                                            <Skeleton className={cn('PlatformIcon')}/> :
                                            platformIcon
                                        }
                                    </div>
                                </div>
                            }
                            {!isCollection &&
                                <div className={cn('liveCountUp')}>
                                    <p className={cn('ItemText')}>
                                        {isLoadingNft ?
                                            <Skeleton /> :
                                            <>
                                                {price === 0 ? 'Free Mint' : `${pr} sol`}
                                            </>
                                        }
                                    </p>
                                    {date.length && <TimeCounter date={date} />}
                                </div>
                            }
                        </div>
                        <div className={cn('RightInfo') + (!isCollection ? ' NotWide' : '')}>
                            {isCollection &&
                                <div className={cn('ProgressBar')}>
                                    <div className={cn('progress')}>
                                        {isLoadingNft && <Skeleton className={cn('skeleton-progress')} />}
                                        <div style={{width: `${isLoadingNft ? 0 : loadingBarWidth}%`, transition: `width ${loadingBarWidth ? 1 : 0}s`}}>
                                            <div className={cn('progressbarinfo')} />
                                        </div>
                                    </div>
                                    <div className={cn('ProgressBarInfo')}>
                                        <div>
                                            {isLoadingNft ?
                                                <Skeleton /> :
                                                <>
                                                    <CountUp duration={1} start={startsMintedAmount} end={mintedAmount} />
                                                    {" / "}
                                                    {supply}
                                                </>
                                            }
                                        </div>
                                        <div>
                                            {isLoadingNft ?
                                                <Skeleton /> :
                                                <>
                                                    <CountUp
                                                        duration={1}
                                                        start={startsPercents}
                                                        end={percents}
                                                    /> %
                                                </>
                                            }
                                        </div>
                                    </div>
                                </div>
                            }
                            {!isCollection &&
                                <div className={cn('PlatformIcon')}>
                                    {isLoadingNft ?
                                        <Skeleton /> :
                                        platformIcon
                                    }
                                </div>
                            }
                        </div>
                    </div>
                </div>
            </li>
        </Flipped>
    )
}
