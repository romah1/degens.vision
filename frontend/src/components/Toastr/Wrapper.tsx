import React, { useCallback } from 'react';
import { Message } from './Message';
import { WrapperProps } from './interfaces';
import { useAppDispatch } from '../../store/hooks';
import { removeToastrMessage } from './redux';

export const DEFAULT_DURATION = 5000;

export const ToastrWrapper: React.FC<WrapperProps> = props => {
    const dispatch = useAppDispatch();

    const onDestroy = useCallback(
        (id: number) => {
            dispatch(removeToastrMessage(id));
        },
        [dispatch],
    );

    return (
        <div
            className={`toastr__wrapper toastr_${props.position}`}
        >
            {
                props.messages.map(message => (
                    <Message
                        key={`toastr_${message.id}`}
                        id={message.id}
                        title={message.title}
                        level={message.level}
                        detail={message.detail}
                        duration={props.duration}
                        onDestroy={onDestroy}
                    />
                ))
            }
        </div>
    );
};
