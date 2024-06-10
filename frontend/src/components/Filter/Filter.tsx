import React, { useCallback, useEffect } from "react";

import block from "bem-cn-lite";
import Slider from "@mui/joy/Slider";
import Input from "@mui/joy/Input";
import Radio from "@mui/joy/Radio";
import RadioGroup from "@mui/joy/RadioGroup";
import List from "@mui/joy/List";
import ListItem from "@mui/joy/ListItem";
import ListItemDecorator from "@mui/joy/ListItemDecorator";
import Checkbox from "@mui/joy/Checkbox";
import Button from "@mui/joy/Button";

import "./Filter.scss";
import { LmnftShapedIcon, TensorShapedIcon, TruffleShapedIcon } from "../icons";

const cn = block("Filter");

const marks = [
    {
        value: 0,
        label: "0",
    },
    {
        value: 100,
        label: "100",
    },
];

const launchpadsData = [
    {
        name: "Launchmynft",
        icon: <LmnftShapedIcon />,
        value: "lmnft",
    },
    {
        name: "Tensor",
        icon: <TensorShapedIcon />,
        value: "tensor",
    },
    {
        name: "Truffle",
        icon: <TruffleShapedIcon />,
        value: "truffle",
    },
];

const defaultLaunchpads = ["lmnft", "tensor", "truffle"];
const onlyFreeMints = false;
const defaultPriceRange = [0, 100];

type Props = {
    className?: string;
    settings: {
        launchpads: string[];
        priceRange: number[];
        onlyFree: boolean;
    };
    onSave: () => void;
    setSettings: React.Dispatch<
        React.SetStateAction<{
            launchpads: string[];
            priceRange: number[];
            onlyFree: boolean;
        }>
    >;
};

export const Filter: React.FC<Props> = ({
    className,
    settings,
    setSettings,
    onSave,
}) => {
    const [launchpads, setLaunchpads] = React.useState<string[]>(
        settings.launchpads
    );
    const [priceRange, setPriceRange] = React.useState<number[]>(
        settings.priceRange
    );
    const [onlyFree, setOnlyFree] = React.useState<boolean>(settings.onlyFree);

    const handleChange = useCallback(
        (event: Event, newValue: number | number[]) => {
            setSettings((prev) => ({
                ...prev,
                priceRange: newValue as number[],
            }));
        },
        []
    );

    const onChange = useCallback(
        (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
            setSettings((prev) => {
                const newValue = [...prev.priceRange];
                newValue[index] = Number(event.target.value);
                return { ...prev, priceRange: newValue };
            });
        },
        []
    );

    const onChangeFreeMints = useCallback(() => {
        setSettings((prev) => ({ ...prev, onlyFree: !prev.onlyFree }));
    }, []);

    const onChangeLauchpads = useCallback(
        (launchpad: string) => (event: React.ChangeEvent<HTMLInputElement>) => {
            if (event.target.checked) {
                setSettings((prev) => ({
                    ...prev,
                    launchpads: [...prev.launchpads, launchpad],
                }));
            } else {
                setSettings((prev) => ({
                    ...prev,
                    launchpads: prev.launchpads.filter(
                        (text) => text !== launchpad
                    ),
                }));
            }
        },
        []
    );

    const resetFilters = useCallback(() => {
        setSettings({
            launchpads: defaultLaunchpads,
            priceRange: defaultPriceRange,
            onlyFree: onlyFreeMints,
        });
    }, []);

    useEffect(() => {
        if (settings.launchpads !== launchpads) {
            setLaunchpads(settings.launchpads);
        }
        if (settings.priceRange !== priceRange) {
            setPriceRange(settings.priceRange);
        }
        if (settings.onlyFree !== onlyFree) {
            setOnlyFree(settings.onlyFree);
        }
    }, [settings]);

    const inputRef = React.useRef<HTMLInputElement | null>(null);

    return (
        <div className={cn() + " " + className}>
            <div className={cn("Block")}>
                <div className={cn("Title")}>Price Range</div>
                <div className={cn("Columns")}>
                    <Input
                        value={priceRange[0]}
                        onChange={onChange(0)}
                        type="number"
                        variant="outlined"
                        slotProps={{
                            input: {
                                ref: inputRef,
                                min: 0,
                                max: 100,
                                step: 0.001,
                            },
                        }}
                    />
                    <Input
                        value={priceRange[1]}
                        onChange={onChange(1)}
                        type="number"
                        variant="outlined"
                        slotProps={{
                            input: {
                                ref: inputRef,
                                min: 0,
                                max: 100,
                                step: 0.001,
                            },
                        }}
                    />
                </div>
                <Slider
                    className={cn("Slider")}
                    value={priceRange}
                    onChange={handleChange}
                    valueLabelDisplay="off"
                    marks={marks}
                    max={100}
                    step={0.001}
                />
                <RadioGroup value={onlyFree ? "yes" : null}>
                    <Radio
                        value="yes"
                        label="Only Free Mints"
                        size="sm"
                        onClick={onChangeFreeMints}
                        className={cn("RadioButton")}
                    />
                </RadioGroup>
            </div>
            <div className={cn("Block")}>
                <div className={cn("Title")}>Launchpads</div>
                <List role="group" orientation="vertical">
                    {launchpadsData.map((item) => (
                        <ListItem
                            key={item.name}
                            className={cn("LaunchpadItem")}
                        >
                            <ListItemDecorator>{item.icon}</ListItemDecorator>
                            <Checkbox
                                disableIcon
                                overlay
                                label={item.name}
                                checked={launchpads.includes(item.value)}
                                onChange={onChangeLauchpads(item.value)}
                            />
                        </ListItem>
                    ))}
                </List>
            </div>
            <div className={cn("ButtonGroup")}>
                <Button
                    variant="outlined"
                    className={cn("Button")}
                    onClick={resetFilters}
                >
                    Reset
                </Button>

                <Button
                    variant="outlined"
                    className={cn("ButtonApply")}
                    onClick={onSave}
                >
                    Apply
                </Button>
            </div>
        </div>
    );
};
