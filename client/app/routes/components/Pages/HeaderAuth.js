import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';


const HeaderAuth = (props) => (
    <div className="mb-4">
        <div className="mb-4 text-center">
            <Link to="/" className="d-inline-block">
                {
                    <i className={`fa fa-${props.icon} fa-3x ${props.iconClassName}`}/>
                }
            </Link>
        </div>
        <h5 className="text-center mb-4">
            { props.title }
        </h5>
    </div>
)
HeaderAuth.propTypes = {
    icon: PropTypes.node,
    iconClassName: PropTypes.node,
    title: PropTypes.node
};
HeaderAuth.defaultProps = {
    title: "Waiting for Data...",
    iconClassName: "text-theme"
};

export { HeaderAuth };
