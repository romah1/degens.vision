import React, { memo, useState, useCallback, useEffect } from "react";

import block from "bem-cn-lite";
import { Button, Tooltip } from "@mui/joy";
import {
    CopiedIcon,
    CopyIcon,
    DiscordIcon,
    HyperSpaceIcon,
    LmnftShapedIcon,
    MEIcon,
    SniperIcon,
    TensorIcon,
    TruffleShapedIcon,
    TwitterIcon,
    ViewIcon,
    WhitepaperIcon,
} from "../icons";

import "./CollectionContainer.scss";
import Skeleton from "react-loading-skeleton";
import CountUp from "react-countup";
import moment from "moment";
import { CopyComponent } from "../Copy/Copy";

const cn = block("CollectionHeader");

interface CollectionHeaderProps {
    image?: string;
    title?: string;
    address?: string;
    mintedAmount?: number;
    supply?: number;
    mintDate?: string;
    mintPrice?: number;
    uniqueMintersCount?: number;
    attributesCount?: number;
    viewsCount?: number;
    collectionType?: string;
    socials?: {
        twitter?: string;
        discord?: string;
        whitepaper?: string;
    };
    mintPage?: string;
    markets?: {
        magiceden?: string;
        sniper?: string;
        tensor?: string;
        hyperspace?: string;
    };
}

const SocialIconMap = {
    twitter: <TwitterIcon />,
    discord: <DiscordIcon />,
    whitepaper: <WhitepaperIcon />,
};

const mintType: Record<string, string> = {
    cnft: 'The collection was created using Bubblegum Standard',
    mplx: 'The collection was created using MPL TM Standard',
    core: 'The collection was created using Metaplex CORE Standard'
}

const MarketsIconMap = {
    tensor: <TensorIcon />,
    magiceden: <MEIcon />,
    hyperspace: <HyperSpaceIcon />,
    sniper: <SniperIcon />
};

export const CollectionHeader: React.FC<CollectionHeaderProps> = memo(
    (props) => {
        const {
            image,
            title,
            address,
            mintedAmount,
            supply,
            mintDate,
            mintPrice,
            uniqueMintersCount,
            attributesCount,
            viewsCount,
            socials,
            mintPage,
            collectionType,
            markets,
        } = props;

        const pr = mintPrice?.toFixed(3) === '0.000' ? '< 0.001' : parseFloat(mintPrice?.toFixed(3) || '0');

        const [imageLoaded, setImageLoaded] = useState(false);

        useEffect(() => {
            if (image === undefined) {
                setImageLoaded(false);
            }
        }, [image]);

        const handleOnLoad = useCallback(() => {
            setImageLoaded(true);
        }, []);

        const handleOnLoadStart = useCallback(() => {
            const timeout = setTimeout(() => setImageLoaded(false), 0);
            return () => clearTimeout(timeout);
        }, []);

        return (
            <div className={cn()}>
                <div className={cn("Header")}>
                    <div className={cn("Image")}>
                        {!imageLoaded && <Skeleton className={cn("Image")} />}
                        {image && (
                            <img
                                src={image}
                                onLoadStart={handleOnLoadStart}
                                onLoad={handleOnLoad}
                                style={{ opacity: imageLoaded ? 1 : 0 }}
                                alt={title}
                                className={cn("Image")}
                            />
                        )}
                    </div>
                    <div className={cn("CollectionInfo")}>
                        <div className={cn("LeftSide")}>
                            <div className={cn("Title")}>
                                <div className={cn("TitleText")}>
                                    {title || <Skeleton />}
                                </div>
                                <div className={cn("Address")}>
                                    {address ? (
                                        <CopyComponent
                                            value={`${address.slice(
                                                0,
                                                10
                                            )}...${address.slice(-8)}`}
                                            copyValue={address}
                                        />
                                    ) : (
                                        <Skeleton />
                                    )}
                                </div>
                            </div>
                            <div className={cn("Socials")}>
                                {mintPage && (
                                    <Button
                                        component="a"
                                        className={cn("MintPageButton")}
                                        variant="plain"
                                        href={mintPage}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        title="Mint Page"
                                    >
                                        Mint Page
                                    </Button>
                                )}

                                {socials &&
                                    Object.entries(socials).map(
                                        ([key, href], index) => {
                                            return (
                                                <Button
                                                    component="a"
                                                    key={index}
                                                    className={cn(
                                                        "ActionButton"
                                                    )}
                                                    title={key}
                                                    target="_blank"
                                                    rel="noopener noreferrer"
                                                    variant="plain"
                                                    startDecorator={
                                                        SocialIconMap?.[
                                                            key as unknown as keyof typeof SocialIconMap
                                                        ]
                                                    }
                                                    href={href}
                                                />
                                            );
                                        }
                                    )}
                                {markets &&
                                    Object.entries(markets).map(
                                        ([key, href], index) => {
                                            return (
                                                <Button
                                                    component="a"
                                                    target="_blank"
                                                    title={key}
                                                    rel="noopener noreferrer"
                                                    key={index}
                                                    className={cn(
                                                        "ActionButtonMarket"
                                                    )}
                                                    variant="plain"
                                                    startDecorator={
                                                        MarketsIconMap?.[
                                                            key as unknown as keyof typeof MarketsIconMap
                                                        ]
                                                    }
                                                    href={href}
                                                />
                                            );
                                        }
                                    )}
                            </div>
                        </div>
                        <div className={cn("RightSide")}>
                            <div className={cn("ViewsCount")}>
                                {viewsCount !== undefined ? (
                                    <ViewIcon />
                                ) : (
                                    <Skeleton />
                                )}
                                {viewsCount !== undefined ? (
                                    <CountUp
                                        start={0}
                                        end={viewsCount}
                                        duration={0.3}
                                    />
                                ) : (
                                    <Skeleton />
                                )}
                            </div>
                            <div className={cn("CollectionType")}>
                                <Tooltip title={mintType?.[collectionType || 'mplx']} className={cn('collectionTypeTooltip')}>
                                    <Button
                                        className={cn("collectionTypeButton")}
                                        variant="plain"
                                    >{collectionType}</Button>
                                </Tooltip>
                            </div>
                        </div>
                    </div>
                </div>
                <div className={cn("OtherInfo")}>
                    <div className={cn("OtherItem")}>
                        <div className={cn("OtherHeader")}>
                            {mintedAmount !== undefined ? (
                                "Minted"
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                        <div className={cn("OtherValue")}>
                            {mintedAmount !== undefined ? (
                                <CountUp
                                    start={0}
                                    end={mintedAmount}
                                    duration={0.3}
                                />
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                    </div>
                    <div className={cn("OtherItem")}>
                        <div className={cn("OtherHeader")}>
                            {supply !== undefined ? "Supply" : <Skeleton />}
                        </div>
                        <div className={cn("OtherValue")}>
                            {supply !== undefined ? (
                                <CountUp
                                    start={0}
                                    end={supply}
                                    duration={0.3}
                                />
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                    </div>
                    <div className={cn("OtherItem")}>
                        <div className={cn("OtherHeader")}>
                            {mintDate ? "Mint Date" : <Skeleton />}
                        </div>
                        <div className={cn("OtherValue")}>
                            {mintDate !== undefined ? (
                                moment(mintDate).format("MMM DD")
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                    </div>
                    <div className={cn("OtherItem")}>
                        <div className={cn("OtherHeader")}>
                            {mintPrice !== undefined ? (
                                "Mint Price"
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                        <div className={cn("OtherValue")}>
                            {mintPrice !== undefined ? (
                                <>
                                    {mintPrice === 0
                                        ? "Free Mint"
                                        : `${pr} SOL`}
                                </>
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                    </div>
                    <div className={cn("OtherItem")}>
                        <div className={cn("OtherHeader")}>
                            {uniqueMintersCount !== undefined ? (
                                "Unique Minters"
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                        <div className={cn("OtherValue")}>
                            {uniqueMintersCount !== undefined ? (
                                <CountUp
                                    start={0}
                                    end={uniqueMintersCount}
                                    duration={0.3}
                                />
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                    </div>
                    <div className={cn("OtherItem")}>
                        <div className={cn("OtherHeader")}>
                            {attributesCount !== undefined ? (
                                "Traits"
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                        <div className={cn("OtherValue")}>
                            {attributesCount !== undefined ? (
                                <CountUp
                                    start={0}
                                    end={attributesCount}
                                    duration={0.3}
                                />
                            ) : (
                                <Skeleton />
                            )}
                        </div>
                    </div>
                </div>
            </div>
        );
    }
);
