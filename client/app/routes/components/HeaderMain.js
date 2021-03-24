import React from 'react';
import PropTypes from 'prop-types';

const HeaderMain = (props) => (
    <React.Fragment>
        <div className={` d-flex ${ props.className }` }>
            <h1 className="display-4 mr-3 mb-0 align-self-start">
                { props.title }
            </h1>
        </div>
    </React.Fragment>
)
HeaderMain.propTypes = {
    title: PropTypes.string,
    subTitle: PropTypes.node,
    className: PropTypes.string
};
HeaderMain.defaultProps = {
    title: "Waiting for Data...",
    subTitle: "Waiting for Data...",
    className: "my-4"
};

export { HeaderMain };
