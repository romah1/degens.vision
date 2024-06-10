import React, { CSSProperties, useCallback, useEffect, useState } from 'react';
import { MessageProps } from './interfaces';
import { DEFAULT_DURATION } from './Wrapper';

import './styles.scss';

export const Message: React.FC<MessageProps> = props => {
    const { duration = DEFAULT_DURATION, detail, title } = props;
    const [timerId, setTimerId] = useState(0);
    const [expanded, setExpanded] = useState(false);
    const [mainStyle, setMainStyle] = useState<CSSProperties>({});

    const makeTimeout = useCallback(
        () => {
            return setTimeout(
                () => {
                    setMainStyle({ opacity: 0 });
                    setTimeout(() => props.onDestroy(props.id), 500);
                },
                duration,
            ) as unknown as number;
        },
        //eslint-disable-next-line
        [duration],
    );

    const handleMouseOver = useCallback(
        () => setTimerId(timerId => {
            clearTimeout(timerId);
            return 0;
        }),
        [setTimerId],
    );

    const handleMouseLeave = useCallback(
        () => setTimerId(makeTimeout()),
        [setTimerId, makeTimeout],
    );

    const handleCloseMessage = useCallback(
        () => props.onDestroy(props.id),
        //eslint-disable-next-line
        [],
    );

    const toggleExpand = useCallback(
        () => setExpanded(expanded => !expanded),
        [setExpanded],
    );

    useEffect(
        () => setTimerId(makeTimeout()),
        [setTimerId, makeTimeout],
    );

    return (
        <div
            className={`toastr__message_${props.level}`}
            style={mainStyle}
            onMouseOver={handleMouseOver}
            onMouseLeave={handleMouseLeave}
        >
            <div
                className="toastr__timer"
                style={timerId ? { animation: `${duration}ms linear timer-indicator`, opacity: 1 } : { opacity: 0 }}
            />
            <div className="toastr__message-field">
                <div className="toastr__message-title">
                    <div className="toastr__message-title_text">{title}</div>
                    <div className="toastr__message-close" onClick={handleCloseMessage}>X</div>
                </div>
                {
                    detail && (
                        <>
                            <div className="toastr__message-detail" style={expanded ? { display: 'block' } : {}}>
                                {detail}
                            </div>
                            <div
                                className="toastr__message-detail_button"
                                onClick={toggleExpand}
                            >
                                {expanded ? 'Collapse' : 'Detailed'}
                            </div>
                        </>
                    )
                }
            </div>
        </div>
    );
};
