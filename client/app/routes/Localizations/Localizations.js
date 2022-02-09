import React from 'react';

import {
    Container,
    Row,
    Col,
} from './../../components';

import {HeaderMain} from '../components/HeaderMain';

import {LocalizationTable} from './Table';
import {Action} from './Action';
import {Media} from 'reactstrap';
import {toast, ToastContainer} from 'react-toastify';
import {connect} from 'react-redux';
import {
    createLocalizationRequest,
    listLocalizationsRequest,
} from '../../redux/localizations/actions';
import {listLanguagesRequest} from '../../redux/languages/actions';

import * as PropTypes from 'prop-types';

const defaultChosenLocalization = {locale: '', name: ''};

class Localizations extends React.Component {
    constructor(props) {
        super(props);

        this.props.listLocalizationsRequest();
        this.props.listLanguagesRequest();
        this.state = {
            chosenLocalization: defaultChosenLocalization,
        };
    }

    handleChooseLocalization = (data) => {
        let localization;
        if (data.length === 0) {
            return;
        }
        localization = this.props.listLanguages.find(o => o.code === data[0].code);

        this.setState(function (previousState) {
            return {...previousState, chosenLocalization: localization};
        });
    };

    handleSubmitLocalization = (e) => {
        e.preventDefault();
        const {code, name} = this.state.chosenLocalization;
        if (code.trim() === '' || name.trim() === '') {
            toast.success(contentSuccess('Value not valid!'));
            return;
        }
        this.props.createLocalizationRequest(code, name);
    };

    render() {
        const {listLocalizations, listLanguages, errorMsg} = this.props;
        if (errorMsg !== '') {
            alert(errorMsg);
        }

        return (
            <Container>
                <HeaderMain
                    title="Localizations"
                    className="mb-5 mt-4"
                />
                <Row>
                    <Col lg={8}>
                        <LocalizationTable
                            listLocalizations={listLocalizations}
                        />
                    </Col>
                    <Col lg={4}>
                        <Action
                            chosenLocalization={this.state.chosenLocalization}
                            handleChooseLocalization={this.handleChooseLocalization}
                            handleSubmitLocalization={this.handleSubmitLocalization}
                            listLanguages={listLanguages}
                        />
                    </Col>
                </Row>
                <ToastContainer
                    position="top-right"
                    autoClose={3000}
                    draggable={false}
                    hideProgressBar={true}
                    limit={3}
                />
            </Container>
        );
    }
}

// eslint-disable-next-line react/prop-types
const contentSuccess = (description) => (
    <Media>
        <Media middle left className="mr-3">
            <i className="fa fa-fw fa-2x fa-check"/>
        </Media>
        <Media body>
            <Media heading tag="h6">
                Success!
            </Media>
            <p>{description}</p>
        </Media>
    </Media>
);

Localizations.propTypes = {
    listLocalizations: PropTypes.array,
    listLanguages: PropTypes.array,
    listLocalizationsRequest: PropTypes.func.isRequired,
    listLanguagesRequest: PropTypes.func.isRequired,
    createLocalizationRequest: PropTypes.func.isRequired,
    errorMsg: PropTypes.string,
};

const mapStateToProps = (state) => ({
    listLocalizations: state.localizations.listLocalizations,
    listLanguages: state.languages.listLanguages,
    errorMsg: state.localizations.error,
});

const mapDispatchToProps = {
    listLocalizationsRequest,
    createLocalizationRequest,
    listLanguagesRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(Localizations);