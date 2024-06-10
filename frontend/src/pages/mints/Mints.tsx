import React, { useEffect, useState, useMemo, memo, useCallback } from "react";

import block from "bem-cn-lite";
import { BarChart } from "@mui/x-charts/BarChart";
import AccordionGroup from "@mui/joy/AccordionGroup";
import { axisClasses } from "@mui/x-charts/ChartsAxis";
import Modal from "@mui/joy/Modal";
import { Transition } from "react-transition-group";
import ModalClose from "@mui/joy/ModalClose";
import { Navigate } from "react-router-dom";

import { Container } from "src/components/Container/Container";
import { ContainerCap } from "src/components/ContainerCap/ContainerCap";
import {
    LmnftShapedIcon,
    TensorShapedIcon,
    PauseIcon,
    RefreshIcon,
    FilterIcon,
    PlayIcon,
} from "src/components/icons";
import {
    useGetCollectionInfoMutation,
    useGetInitLiveFeedQuery,
} from "src/features/collections/collectionsApi";
import { useErrorToaster } from "src/components";
import { CopyComponent } from "src/components/Copy/Copy";

import "./Mints.scss";
import { NFTItem } from "src/components/NFTItem/NFTItem";
import { CollectionContainer } from "src/components/CollectionContainer/CollectionContainer";
import { CollectionHeader } from "src/components/CollectionContainer/CollectionHeader";
import { CollectionBlock } from "src/components/CollectionContainer/CollectionBlock";
import { ContainerBody } from "src/components/ContainerBody/ContainerBody";
import { Cards } from "src/components/Cards/Cards";
import {
    CollectionInfoResponse,
    LiveMintsItem,
} from "src/interfaces/collections";
import { Filter } from "src/components/Filter/Filter";

import { reconnectingSocket } from "./utils";

import { TopMintsBlock } from "./TopMints";

const cn = block("MintsPage");

const loadingItems = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
const loadingBar = [
    1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
    22, 23, 24, 25,
];

type MintInfoBlockProps = {
    selectedCollection: string | null;
};

const valueFormatter = (value: number) => {
    if (value === 26) {
        return `>25`;
    }

    return value.toString();
};

const chartSetting = {
    series: [{ dataKey: "holders_amount" }],
    height: 350,
    axisHighlight: {
        x: "none" as "band" | "none" | "line",
        y: "none" as "none" | "line",
    },
    grid: { horizontal: true },
    margin: { left: 50, right: 0, bottom: 25, top: 15 },

    sx: {
        [`& .${axisClasses.directionY} .${axisClasses.label}`]: {
            transform: "translate(60px, -120px)",
        },
        [`& .${axisClasses.directionY} .${axisClasses.label} text`]: {
            transform: "rotate(0deg)",
            color: "#fff",
        },

        [`& .${axisClasses.directionX} .${axisClasses.label}`]: {
            transform: "translate(40%, 0px)",
            color: "#fff",
        },
    },
};

const MintInfoBlock: React.FC<MintInfoBlockProps> = memo(
    ({ selectedCollection }) => {
        const [hasData, setHasData] = useState(true);

        const [memoData, setMemoData] = useState<CollectionInfoResponse | null>(
            null
        );

        const [
            useCollectionInfoQuery,
            { data, isLoading: isLoading, error, ...other },
        ] = useGetCollectionInfoMutation();

        const collectData = useCallback(() => {
            if (selectedCollection) {
                useCollectionInfoQuery(selectedCollection);
            }
        }, [useCollectionInfoQuery, selectedCollection]);

        useEffect(() => {
            const ptr = setInterval(collectData, 5000);
            return () => clearInterval(ptr);
        }, [collectData]);

        useEffect(() => {
            setMemoData(null);
        }, [selectedCollection]);

        useErrorToaster(error);

        useEffect(() => {
            if (!selectedCollection) {
                setHasData(false);
            } else {
                setHasData(true);
                useCollectionInfoQuery(selectedCollection);
            }
        }, [selectedCollection]);

        useEffect(() => {
            if (data && JSON.stringify(data) !== JSON.stringify(memoData)) {
                setMemoData(data);
            }
        }, [data, memoData]);

        const defaultData = React.useMemo(() => {
            return loadingBar.map((item) => ({
                holders_amount: 0,
                mints_amount: item,
            }));
        }, []);

        const [showData, totalSum] = React.useMemo(() => {
            if (
                memoData &&
                memoData.mints_distribution &&
                memoData.mints_distribution.length
            ) {
                const response: CollectionInfoResponse["mints_distribution"] =
                    [];
                memoData.mints_distribution.forEach((item) => {
                    if (item.mints_amount >= 26) {
                        if (response.find((item) => item.mints_amount === 26)) {
                            response.find(
                                (item) => item.mints_amount === 26
                            )!.holders_amount += item.holders_amount;
                        } else {
                            response.push({
                                holders_amount: item.holders_amount,
                                mints_amount: 26,
                            });
                        }
                    } else {
                        response.push({
                            holders_amount: item.holders_amount,
                            mints_amount: item.mints_amount,
                        });
                    }
                });

                loadingBar.forEach((loadItem) => {
                    if (
                        !response.find(
                            (item) => item.mints_amount === loadItem
                        ) &&
                        response.length < 25
                    ) {
                        response.push({
                            holders_amount: 0,
                            mints_amount: loadItem,
                        });
                    }
                });

                const newResponse = response.sort(
                    (a, b) => a.mints_amount - b.mints_amount
                );

                const totalSum = Math.max(
                    ...newResponse.map((item) => item.holders_amount)
                );

                return [newResponse, totalSum < 5 ? 5 : totalSum];
            }
            return [defaultData, 5];
        }, [memoData]);

        const header = useMemo(() => {
            if (!hasData)
                return <div className={cn("loader")}>Probably nothing :(</div>;
            const socials: {
                twitter?: string;
                discord?: string;
                whitepaper?: string;
            } = {};
            const mintPage: Record<string, string> = {};
            const markets: {
                magiceden?: string;
                tensor?: string;
                sniper?: string;
                hyperspace?: string;
            } = {};

            for (const item of memoData?.collection.links || []) {
                if (item.type === "mint") {
                    mintPage.uri = item.uri;
                } else {
                    socials[item.type as "twitter" | "discord" | "whitepaper"] =
                        item.uri;
                }
            }

            for (const item of memoData?.collection.market_links || []) {
                markets[
                    item.type as
                        | "magiceden"
                        | "tensor"
                        | "sniper"
                        | "hyperspace"
                ] = item.uri;
            }

            return (
                <CollectionHeader
                    image={memoData?.collection.image}
                    title={
                        memoData?.collection.name || memoData?.collection.symbol
                    }
                    address={memoData?.collection.collection_key}
                    mintedAmount={memoData?.collection.mints_count}
                    supply={memoData?.collection.size}
                    mintDate={memoData?.collection?.date}
                    mintPrice={memoData?.collection.mint_price}
                    collectionType={memoData?.collection.type}
                    uniqueMintersCount={
                        memoData?.collection.unique_minters_amount
                    }
                    attributesCount={memoData?.collection.attributes_amount}
                    viewsCount={undefined}
                    mintPage={mintPage?.uri}
                    socials={socials}
                    markets={markets}
                />
            );
        }, [hasData, memoData]);

        const bar = useMemo(() => {
            if (!hasData) return null;

            return (
                <AccordionGroup>
                    <CollectionBlock
                        name="Unique Minters Distribution"
                        secondaryText="Holders Distribution"
                    >
                        <div className={cn("ChartLabel")}>Unique Minters</div>
                        <BarChart
                            yAxis={[
                                {
                                    tickMinStep: 1,
                                    max: totalSum,
                                },
                            ]}
                            dataset={showData}
                            xAxis={[
                                {
                                    scaleType: "band",
                                    dataKey: "mints_amount",
                                    label: "NFTs Quantity",
                                    max: 50,
                                    valueFormatter: valueFormatter,
                                },
                            ]}
                            {...chartSetting}
                        />
                        <div
                            className={cn("ChartLabel")}
                            style={{ textAlign: "right" }}
                        >
                            NFTs Quantity
                        </div>
                    </CollectionBlock>
                </AccordionGroup>
            );
        }, [hasData, totalSum, showData]);

        return (
            <CollectionContainer>
                {header}
                {bar}
            </CollectionContainer>
        );
    }
);

interface LiveMintsBlockProps {
    setSelectedCollection: React.Dispatch<React.SetStateAction<string | null>>;
}

const LiveMintsBlock: React.FC<LiveMintsBlockProps> = memo(
    ({ setSelectedCollection }) => {
        const { data: initData, isLoading } = useGetInitLiveFeedQuery();
        const [messages, setMessages] = useState<LiveMintsItem[]>([]);

        useEffect(() => {
            function handleMessage(message: LiveMintsItem) {
                setMessages([message, ...messages]);
            }
            websocketClient.on(handleMessage);
            return () => websocketClient.off(handleMessage);
        }, [messages, setMessages]);

        const websocketClient = useMemo(
            () => reconnectingSocket(Date.now()),
            []
        );

        const [open, setOpen] = useState<boolean>(false);
        const [modalData, setModalData] = useState<LiveMintsItem | null>(null);

        const [isConnected, setIsConnected] = useState(
            websocketClient.isConnected()
        );

        useEffect(() => {
            return websocketClient.onStateChange(setIsConnected);
        }, [setIsConnected]);

        useEffect(() => {
            return () => websocketClient.close();
        }, []);

        const handleNftClick = useCallback(
            (nftData: LiveMintsItem) => (e: React.MouseEvent<HTMLElement>) => {
                e.stopPropagation();
                setModalData(nftData);
                setOpen(true);
            },
            []
        );

        const handleCloseModal = useCallback(() => {
            setOpen(false);
        }, []);

        const [items, setItems] = useState<LiveMintsItem[]>([]);
        const [isActive, setIsActive] = useState(true);
        const [isFilterOpen, setIsFilterOpen] = useState(false);
        const [settings, setSettings] = useState<{
            launchpads: string[];
            priceRange: number[];
            onlyFree: boolean;
        }>({
            launchpads: ["lmnft", "tensor", "truffle"],
            priceRange: [0, 100],
            onlyFree: false,
        });

        useEffect(() => {
            if (!isActive || !messages.length) {
                return;
            }
            const data = messages[0];
            if (
                !settings.launchpads.includes(data.collection.launchpad) ||
                (data.mint.mint_price !== 0 && settings.onlyFree) ||
                (!settings.onlyFree &&
                    (data.mint.mint_price < settings.priceRange[0] ||
                        data.mint.mint_price > settings.priceRange[1]))
            ) {
                return;
            }

            setItems((prev) => {
                const newItems = [data, ...prev];
                if (newItems.length > 20) return newItems.slice(0, 20);
                return newItems;
            });
        }, [messages]);

        useEffect(() => {
            if (initData) {
                setItems(initData);
            }
        }, [initData]);

        const handleChooseCollection = useCallback(
            (collection_key: string | undefined) => () => {
                if (collection_key) {
                    setSelectedCollection(collection_key);
                    setOpen(false);
                }
            },
            []
        );

        const handleIconClick = useCallback(() => {
            setIsActive((prev) => !prev);
        }, []);

        const handleFilterClick = useCallback(() => {
            setIsFilterOpen((prev) => !prev);
        }, []);

        const [flipKey, data] = React.useMemo(() => {
            if (isLoading) {
                return [
                    loadingItems.join(""),
                    loadingItems.map((item) => (
                        <NFTItem
                            key={item.toString()}
                            collectionKey={item.toString()}
                            isLoading={true}
                            name={"Loading..."}
                            price={0}
                            supply={0}
                            startsMintsAmount={0}
                            startsMintedAmount={0}
                            startsPercents={0}
                            img={""}
                            date={""}
                            platformIcon={<LmnftShapedIcon />}
                        />
                    )),
                ];
            }
            const filteredItems = items.filter((data) => {
                if (
                    !settings.launchpads.includes(data.collection.launchpad) ||
                    (data.mint.mint_price !== 0 && settings.onlyFree) ||
                    (!settings.onlyFree &&
                        (data.mint.mint_price < settings.priceRange[0] ||
                            data.mint.mint_price > settings.priceRange[1]))
                ) {
                    return false;
                }
                return true;
            });

            return [
                filteredItems.map((item) => item.mint.tx_signature).join(""),
                filteredItems.map((item) => (
                    <NFTItem
                        key={item.mint.tx_signature}
                        collectionKey={item.mint.tx_signature}
                        name={item.mint.asset_name || item.collection.name}
                        price={item.mint.mint_price}
                        supply={0}
                        startsMintsAmount={0}
                        startsMintedAmount={0}
                        startsPercents={0}
                        date={item.mint.block_time}
                        img={item.mint.asset_image}
                        platformIcon={
                            item.collection.launchpad === "lmnft" ? (
                                <LmnftShapedIcon />
                            ) : (
                                <TensorShapedIcon />
                            )
                        }
                        onClick={handleChooseCollection(
                            item.collection.collection_key
                        )}
                        onIconClick={handleNftClick(item)}
                    />
                )),
            ];
        }, [items, initData, isLoading, settings]);

        return (
            <Container>
                <ContainerCap
                    name={"Live Mints"}
                    actionButton={<FilterIcon onClick={handleFilterClick} />}
                    nameActionButton={
                        isActive ? (
                            <PauseIcon onClick={handleIconClick} />
                        ) : (
                            <PlayIcon onClick={handleIconClick} />
                        )
                    }
                />
                <Filter
                    className={isFilterOpen ? "Active" : ""}
                    settings={settings}
                    setSettings={setSettings}
                    onSave={handleFilterClick}
                />
                <ContainerBody key={"live-mints"} flipKey={flipKey}>
                    {data}
                </ContainerBody>
                <Transition in={open} timeout={300}>
                    {(state: string) => (
                        <Modal
                            keepMounted
                            open={!["exited", "exiting"].includes(state)}
                            onClose={handleCloseModal}
                            className={cn("ModalOuter")}
                            slotProps={{
                                backdrop: {
                                    sx: {
                                        opacity: 0,
                                        backdropFilter: "none",
                                        transition: `opacity 300ms, backdrop-filter 300ms`,
                                        ...{
                                            entering: {
                                                opacity: 1,
                                                backdropFilter: "blur(8px)",
                                            },
                                            entered: {
                                                opacity: 1,
                                                backdropFilter: "blur(8px)",
                                            },
                                        }[state],
                                    },
                                },
                            }}
                            sx={{
                                visibility:
                                    state === "exited" ? "hidden" : "visible",
                            }}
                        >
                            <div
                                className={cn("ModalInner")}
                                style={{
                                    opacity: ["entering", "entered"].includes(
                                        state
                                    )
                                        ? 1
                                        : 0,
                                }}
                            >
                                <ModalClose variant="plain" sx={{ m: 1 }} />
                                <div className={cn("LeftModalInfo")}>
                                    <img src={modalData?.mint.asset_image} />
                                    <div className={cn("LeftModalInfoBlock")}>
                                        <div className="info-row">
                                            <div className="row-header">
                                                Mint Adress
                                            </div>
                                            <div className="row-body">
                                                <CopyComponent
                                                    value={`${modalData?.mint.asset_key.slice(
                                                        0,
                                                        6
                                                    )}...${modalData?.mint.asset_key.slice(
                                                        -4
                                                    )}`}
                                                    copyValue={
                                                        modalData?.mint
                                                            .asset_key
                                                    }
                                                />
                                            </div>
                                        </div>
                                        <div className="info-row">
                                            <div className="row-header">
                                                Collection Key
                                            </div>
                                            <div className="row-body">
                                                <CopyComponent
                                                    value={`${modalData?.mint.collection_key.slice(
                                                        0,
                                                        6
                                                    )}...${modalData?.mint.collection_key.slice(
                                                        -4
                                                    )}`}
                                                    copyValue={
                                                        modalData?.mint
                                                            .collection_key
                                                    }
                                                />
                                            </div>
                                        </div>
                                        <div className="info-row">
                                            <div className="row-header">
                                                Signature
                                            </div>
                                            <div className="row-body">
                                                <CopyComponent
                                                    value={`${modalData?.mint.tx_signature.slice(
                                                        0,
                                                        6
                                                    )}...${modalData?.mint.tx_signature.slice(
                                                        -4
                                                    )}`}
                                                    copyValue={
                                                        modalData?.mint
                                                            .tx_signature
                                                    }
                                                />
                                            </div>
                                        </div>
                                        {/* <div className='info-row'>
                                        <div className='row-header'>Owner</div>
                                        <div className='row-body'>
                                            <CopyComponent
                                                value={`${modalData?.mint.asset_owner_key.slice(0, 6)}...${modalData?.mint.asset_owner_key.slice(-4)}`}
                                                copyValue={modalData?.mint.asset_owner_key}
                                            />
                                        </div>
                                    </div> */}
                                        {/* <div className='info-row'>
                                        <div className='row-header'>Royalty</div>
                                        <div className='row-body'>0 %</div>
                                    </div> */}
                                    </div>
                                </div>
                                <div className={cn("RightModalInfoBlock")}>
                                    <div className="header">
                                        <div
                                            className="collection-name"
                                            onClick={handleChooseCollection(
                                                modalData?.collection
                                                    .collection_key
                                            )}
                                        >
                                            {modalData?.collection.name || modalData?.collection.symbol}
                                        </div>
                                        <div className="nft-name">
                                            {modalData?.mint.asset_name}
                                        </div>
                                        <div className="owner-name">
                                            <div className="text-before">
                                                Owned by
                                            </div>
                                            <CopyComponent
                                                className="text-after"
                                                value={`${modalData?.mint.asset_owner_key.slice(
                                                    0,
                                                    6
                                                )}...${modalData?.mint.asset_owner_key.slice(
                                                    -4
                                                )}`}
                                                copyValue={
                                                    modalData?.mint
                                                        .asset_owner_key
                                                }
                                            />
                                        </div>
                                    </div>
                                    <div className="description">
                                        <div className="description-header">
                                            Description
                                        </div>
                                        <div className="description-content">
                                            {modalData?.mint.asset_description.length ? modalData?.mint.asset_description : 'Probably nothing :('}
                                        </div>
                                    </div>
                                    <div className="traits">
                                        <div className="traits-header">
                                            Traits
                                        </div>
                                        <div className="traits-content">
                                            <Cards cards={modalData?.mint.asset_attributes || []} />
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </Modal>
                    )}
                </Transition>
            </Container>
        );
    }
);

export const MintsPage: React.FC = () => {
    const [selectedCollection, setSelectedCollection] = useState<string | null>(
        localStorage.getItem("collectionHash")
    );

    const sessionToken = localStorage.getItem("token");

    if (!sessionToken) {
        return <Navigate to="/" />;
    }

    useEffect(() => {
        const collection = localStorage.getItem("collectionHash");

        if (selectedCollection && collection !== selectedCollection) {
            localStorage.setItem("collectionHash", selectedCollection);
        }
    }, [selectedCollection]);

    return (
        <div className={cn()}>
            <TopMintsBlock setSelectedCollection={setSelectedCollection} />
            <MintInfoBlock selectedCollection={selectedCollection} />
            <LiveMintsBlock setSelectedCollection={setSelectedCollection} />
        </div>
    );
};
