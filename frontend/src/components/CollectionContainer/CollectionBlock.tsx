import React, { useState, memo } from "react";

import Accordion from '@mui/joy/Accordion';
import AccordionDetails from '@mui/joy/AccordionDetails';
import AccordionSummary from '@mui/joy/AccordionSummary';

import block from 'bem-cn-lite';

import './CollectionContainer.scss';
import { ChevronDownIcon, ChevronUpIcon } from "../icons";

const cn = block('CollectionBlock');

interface CollectionBlockProps {
    name: string;
    secondaryText?: string;
    isOpen?: boolean;
    children: React.ReactNode;
}

export const CollectionBlock: React.FC<CollectionBlockProps> = memo((props) => {
    return (

        <Accordion className={cn()} defaultExpanded={true}>
            <AccordionSummary className={cn('Cap')}>
                <div className={cn('CapHeader')}>
                    {props.name}
                    {props.secondaryText && (
                        <>
                            <div className={cn('Delimeter')}>/</div>
                            <div className={cn('Disabled')}>{props.secondaryText}</div>
                        </>
                    )}
                </div>
            </AccordionSummary>
            <AccordionDetails className={cn('Details')}>
                {props.children}
            </AccordionDetails>
        </Accordion>

        // <div className={cn()}>
        //     <div className={cn('Cap')} onClick={() => setIsOpen(o => !o)}>

        //         {isOpen ? <ChevronUpIcon /> : <ChevronDownIcon />}
        //     </div>
        //     {isOpen && (
        //         <div className={cn('Body')}>

        //         </div>
        //     )}
        // </div>
    )
})
