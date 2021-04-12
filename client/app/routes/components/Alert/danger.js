import {Alert} from "reactstrap";
import React from "react";
import * as PropTypes from "prop-types";

export const DangerAlert = ({text}) => (
    <Alert color="danger">
        <i className="fa fa-times-circle mr-1 alert-icon"/>
        <span>
            <strong className="alert-heading">Danger! </strong>
            {text}
        </span>
    </Alert>
)

DangerAlert.propTypes = {
    text: PropTypes.string.isRequired,
};
