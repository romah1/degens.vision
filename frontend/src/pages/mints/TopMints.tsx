import React, { useEffect, useState, useCallback, memo, useMemo } from "react";

import { Container } from "src/components/Container/Container";
import { ContainerCap } from "src/components/ContainerCap/ContainerCap";
import { ContainerBody } from "src/components/ContainerBody/ContainerBody";
import { useErrorToaster } from "src/components";
import { NFTItem } from "src/components/NFTItem/NFTItem";

import block from "bem-cn-lite";
import { Button, ToggleButtonGroup } from "@mui/joy";

import { LmnftShapedIcon, RefreshIcon, TensorShapedIcon } from "src/components/icons";

import { useGetCollectionsTopMutation } from "src/features/collections/collectionsApi";

import { CollectionInfo } from "src/interfaces/collections";

const cn = block("MintsPage");

const times = [
    "1 min",
    "5 min",
    "10 min",
    "30 min",
    "1 h",
    "3 h",
    "12 h",
    "24 h",
];
const timesMap: Record<string, number> = {
    "1 min": 1,
    "5 min": 5,
    "10 min": 10,
    "30 min": 30,
    "1 h": 60,
    "3 h": 60 * 3,
    "12 h": 60 * 12,
    "24 h": 60 * 24,
};

const loadingItems = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

type LoadDataProps = {
    flipKey: string;
    mintsAmount: string;
    mintedAmount: string;
    components: CollectionInfo[];
    nftItems?: React.ReactNode[];
};

type CapProps = {
    selectedTime: string;
    setSelectedTime: React.Dispatch<React.SetStateAction<string>>;
};

const Cap: React.FC<CapProps> = memo(({ selectedTime, setSelectedTime }) => {
    const changeToggleButton = useCallback(
        (
            _: React.MouseEvent<HTMLElement, MouseEvent>,
            newValue: string | null
        ) => {
            if (newValue === null) return;

            setSelectedTime(newValue);
        },
        []
    );

    const timesButtons = useMemo(() => {
        return times.map((time, index) => (
            <Button
                className={cn("TimesButton")}
                value={time}
                key={"time_" + index}
            >
                {time}
            </Button>
        ));
    }, []);

    return (
        <ContainerCap name={"Top Mints"}>
            <ToggleButtonGroup
                value={selectedTime}
                onChange={changeToggleButton}
                className={cn("TimesButtonGroup")}
            >
                {timesButtons}
            </ToggleButtonGroup>
        </ContainerCap>
    );
});

interface BodyProps extends TopMintsBlockProps {
    selectedTime: string;
}

const Body: React.FC<BodyProps> = memo(
    ({ selectedTime, setSelectedCollection }) => {
        const [loadedData, setLoadedData] = useState<
            Record<string, LoadDataProps>
        >({});
        const [currentData, setCurrentData] = useState<LoadDataProps>({
            flipKey: "",
            mintsAmount: "",
            mintedAmount: "",
            components: [],
        });
        const [useCollectionsTopQuery, { data, error }] =
            useGetCollectionsTopMutation();

        const collectData = useCallback(() => {
            useCollectionsTopQuery(timesMap[selectedTime] || 1);
        }, [useCollectionsTopQuery, selectedTime]);

        useEffect(() => {
            const ptr = setInterval(collectData, 5000);
            return () => clearInterval(ptr);
        }, [collectData]);

        useEffect(() => {
            const time = localStorage.getItem("topMintsTimeRange");

            if (selectedTime && time !== selectedTime) {
                localStorage.setItem("topMintsTimeRange", selectedTime);
            }
            useCollectionsTopQuery(timesMap[selectedTime] || 1);
        }, [selectedTime]);

        useErrorToaster(error);

        const selectCollection = useCallback(
            (collectionHash: string) => () => {
                setSelectedCollection((oldCollectionHash) =>
                    oldCollectionHash === collectionHash
                        ? oldCollectionHash
                        : collectionHash
                );
            },
            []
        );

        useEffect(() => {
            if (data && data.collections) {
                const collections = data.collections;
                if (collections.length > 0) {
                    const collectionHash = collections[0].collection_key;
                    setSelectedCollection((oldCollectionHash) =>
                        oldCollectionHash ? oldCollectionHash : collectionHash
                    );
                }

                const flipKey = collections
                    .map((item) => item.collection_key)
                    .join("");
                setLoadedData((oldData) => ({
                    ...oldData,
                    [selectedTime]: {
                        flipKey: flipKey.length === 0 ? "blank" : flipKey,
                        mintsAmount: collections
                            .map((item) => item.mints_count_period)
                            .join(""),
                        mintedAmount: collections
                            .map((item) => item.mints_count_total)
                            .join(""),
                        components: collections,
                    },
                }));
            }
        }, [data]);

        const [loadingFlipKey, loadingItemsData] = React.useMemo(() => {
            return [
                loadingItems.map((item) => item.toString()).join(""),
                loadingItems.map((item) => (
                    <NFTItem
                        key={item.toString()}
                        collectionKey={item.toString()}
                        isLoading
                        date=""
                        name="Purple Purple"
                        price={0}
                        mintsAmount={0}
                        mintedAmount={0}
                        supply={0}
                        startsMintsAmount={0}
                        startsMintedAmount={0}
                        startsPercents={0}
                        img=""
                        platformIcon={<LmnftShapedIcon />}
                    />
                )),
            ];
        }, [loadingItems]);

        useEffect(() => {
            if (loadedData[selectedTime]) {
                const data = loadedData[selectedTime];
                if (
                    data.flipKey !== currentData.flipKey ||
                    data.mintedAmount !== currentData.mintedAmount ||
                    data.mintsAmount !== currentData.mintsAmount
                ) {
                    setCurrentData((prevData) => {
                        const newData = { ...data };
                        const nftItems: React.ReactNode[] = [];

                        for (const component of data.components) {
                            const item = prevData.components.find(
                                (item) =>
                                    item.collection_key ===
                                    component.collection_key
                            );
                            let startsMintsAmount = 0;
                            let startsMintedAmount = 0;
                            let startsPercents = 0;
                            if (item) {
                                startsMintsAmount = item.mints_count_period;
                                startsMintedAmount =
                                    item.mints_count_total || 0;
                                startsPercents = Math.floor(
                                    (startsMintedAmount / item.size) * 100
                                );
                            }
                            nftItems.push(
                                <NFTItem
                                    key={component.collection_key}
                                    collectionKey={component.collection_key}
                                    isLoading={false}
                                    name={component.name || component.symbol}
                                    price={component.mint_price}
                                    mintsAmount={component.mints_count_period}
                                    mintedAmount={component.mints_count_total}
                                    supply={component.size}
                                    img={component.image}
                                    date=""
                                    platformIcon={
                                        component.launchpad === "lmnft" ? (
                                            <LmnftShapedIcon />
                                        ) : (
                                            <TensorShapedIcon />
                                        )
                                    }
                                    onClick={selectCollection(
                                        component.collection_key
                                    )}
                                    startsMintsAmount={startsMintsAmount}
                                    startsMintedAmount={startsMintedAmount}
                                    startsPercents={startsPercents}
                                    percents={Math.floor(
                                        ((component.mints_count_total || 0) /
                                            component.size) *
                                            100
                                    )}
                                />
                            );
                        }
                        newData.nftItems = nftItems;

                        return newData;
                    });
                }
            } else {
                if (currentData.flipKey !== "") {
                    setCurrentData({
                        flipKey: "",
                        mintsAmount: "",
                        mintedAmount: "",
                        components: [],
                    });
                }
            }
        }, [loadedData, selectedTime]);

        const isLoadingData = React.useMemo(() => {
            if (currentData.flipKey.length === 0) {
                return true;
            }
            return false;
        }, [currentData]);

        const [flipKey, items] = React.useMemo(() => {
            if (isLoadingData) {
                return [loadingFlipKey, loadingItemsData];
            }
            return [currentData.flipKey, currentData.nftItems];
        }, [isLoadingData, currentData, loadingFlipKey, loadingItemsData]);

        return (
            <ContainerBody key={"container"} flipKey={flipKey}>
                {items}
            </ContainerBody>
        );
    }
);

interface TopMintsBlockProps {
    setSelectedCollection: React.Dispatch<React.SetStateAction<string | null>>;
}

export const TopMintsBlock: React.FC<TopMintsBlockProps> = ({
    setSelectedCollection,
}) => {
    const [selectedTime, setSelectedTime] = useState<string>(
        localStorage.getItem("topMintsTimeRange") || times[0]
    );

    return (
        <Container>
            <Cap
                selectedTime={selectedTime}
                setSelectedTime={setSelectedTime}
            />
            <Body
                selectedTime={selectedTime}
                setSelectedCollection={setSelectedCollection}
            />
        </Container>
    );
};
