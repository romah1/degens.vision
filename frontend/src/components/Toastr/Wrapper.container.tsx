import React from 'react';
import { WrapperContainerProps } from './interfaces';
import { ToastrWrapper as BaseComponent } from './Wrapper';
import { WrapperHook } from './Wrapper.hook';

export const ToastrWrapper: React.FC<WrapperContainerProps> = props => {
    const { messages } = WrapperHook();

    return (
        <BaseComponent
            position={props.position}
            messages={messages}
            duration={props.duration}
        />
    );
};
