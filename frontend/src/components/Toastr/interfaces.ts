import { ReactNode } from 'react';

export type ToastrLevel = 'info' | 'warning' | 'error' | 'success';
export type ToastrPosition = 'bottom-right' | 'bottom-left' | 'top-right' | 'top-left';

export interface Message {
    id: number;
    title: string;
    detail?: ReactNode;
    level: ToastrLevel;
}

export interface MessageProps extends Message{
    onDestroy: (id: number) => void;
    duration?: number;
}

export interface WrapperProps {
    position: ToastrPosition;
    messages: Message[];
    duration?: number;
}

export type WrapperContainerProps = Omit<WrapperProps, 'messages'>;
