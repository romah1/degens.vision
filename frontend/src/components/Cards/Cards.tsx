import React, { useCallback, useEffect, useMemo } from "react";

import block from "bem-cn-lite";

import "./Cards.scss";

const cn = block("Cards");

type Trait = {
    value: string;
    trait_type: string;
}

type CardsProps = {
    cards: Trait[];
    classname?: string;
    traitClassname?: string;
}

type CardProps = {
    card: Trait;
    classname?: string;
    // onClick?: (card: Record<string, string>) => void;
}

const Card: React.FC<CardProps> = (props) => {
    return (
        <div className={cn("card")}>
            <div className={cn("cardTitle")}>
                {props.card.trait_type}
            </div>
            <div className={cn("cardBody")}>
                {props.card.value}
            </div>
        </div>
    )
}


export const Cards: React.FC<CardsProps> = (props) => {
    const {
        cards,
        classname,
        traitClassname
    } = props;

    const traitsList = useMemo(() => {
        return cards.map((card, index) => {
            return <Card key={index} card={card} classname={traitClassname} />
        });
    }, [cards, traitClassname]);

    return (
        <div className={cn() + ' ' + classname}>
            {traitsList.length ? traitsList : 'Probably nothing :('}
        </div>
    )
}
