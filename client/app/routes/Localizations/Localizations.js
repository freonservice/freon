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
import * as PropTypes from 'prop-types';
import lang from '../../../lang.json';

const defaultChosenLocalization = {locale: '', name: ''};

class Localizations extends React.Component {
    constructor(props) {
        super(props);

        this.props.listLocalizationsRequest();
        this.state = {
            languages: JSON.parse(JSON.stringify(lang)),
            chosenLocalization: defaultChosenLocalization,
        };
    }

    handleChooseLocalization = (data) => {
        let localization;
        if (data.length === 0) {
            return;
        }
        localization = this.state.languages.find(o => o.name === data[0].name);

        this.setState(function (previousState) {
            return {...previousState, chosenLocalization: localization};
        });
    };

    handleSubmitLocalization = (e) => {
        e.preventDefault();
        const {locale, name} = this.state.chosenLocalization;
        if (locale.trim() === '' || name.trim() === '') {
            toast.success(contentSuccess('Value not valid!'));
            return;
        }
        this.props.createLocalizationRequest(locale, name);
        // toast.success(contentSuccess('Localization successful added!'));
    };

    render() {
        const {listLocalizations, errorMsg} = this.props;
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
                            languages={this.state.languages}
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
    listLocalizationsRequest: PropTypes.func.isRequired,
    createLocalizationRequest: PropTypes.func.isRequired,
    errorMsg: PropTypes.string,
};

const mapStateToProps = (state) => ({
    listLocalizations: state.localizations.listLocalizations,
    errorMsg: state.localizations.error,
});

const mapDispatchToProps = {
    listLocalizationsRequest,
    createLocalizationRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(Localizations);