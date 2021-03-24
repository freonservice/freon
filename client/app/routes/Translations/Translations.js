import React from 'react';

import {
    Container,
    Row,
    Col,
} from './../../components';

import {HeaderMain} from '../components/HeaderMain';
import {TranslationTable} from './Table';
import {withRouter} from 'react-router-dom';
import * as PropTypes from 'prop-types';
import {connect} from 'react-redux';
import {
    listTranslationsRequest,
    createTranslationRequest,
    updateTranslationRequest,
    hideTranslationRequest,
} from '../../redux/translations/actions';
import {
    listLocalizationsRequest,
} from '../../redux/localizations/actions';

const defaultLocalization = {};

class Translations extends React.Component {
    constructor(props) {
        super(props);

        this.props.listLocalizationsRequest();
        this.props.listTranslationsRequest();

        this.state = {
            chooseLocalization: this.props.listLocalizations.slice(0, 1).shift(),
        };
    }

    handleChooseLocalization = (data) => {
        let localization;
        if (data.length > 0) {
            localization = this.props.listLocalizations.find(o => o.locale === data[0].locale);
        } else {
            localization = defaultLocalization;
        }
        this.setState(function (previousState) {
            return {
                ...previousState,
                chooseLocalization: localization
            };
        });
        this.props.listTranslationsRequest(localization.id || 0);
    };

    handleHideTranslation = (e, id) => {
        e.stopPropagation();
        const translation = this.props.listTranslations.find(o => o.id === id);
        const hide = translation.status === "Active"
        this.props.hideTranslationRequest(translation.id, hide);
    };

    render() {
        const {history, listTranslations, listLocalizations} = this.props;
        return (
            <Container>
                <HeaderMain
                    title="Translations"
                    className="mb-5 mt-4"
                />
                <Row>
                    <Col lg={12}>
                        <TranslationTable
                            listTranslations={listTranslations}
                            listLocalizations={listLocalizations}
                            history={history}
                            chooseLocalization={this.state.chooseLocalization}
                            handleChosenLocalization={this.handleChooseLocalization}
                            handleHideTranslation={this.handleHideTranslation}
                        />
                    </Col>
                </Row>
            </Container>
        );
    }
}

Translations.propTypes = {
    listTranslations: PropTypes.array,
    listLocalizations: PropTypes.array,
    listTranslationsRequest: PropTypes.func.isRequired,
    createTranslationRequest: PropTypes.func.isRequired,
    updateTranslationRequest: PropTypes.func.isRequired,
    hideTranslationRequest: PropTypes.func.isRequired,
    listLocalizationsRequest: PropTypes.func.isRequired,
    errorMsg: PropTypes.string,
    history: PropTypes.shape({
        push: PropTypes.func.isRequired,
    }).isRequired,
};

Translations.defaultProps = {
    listLocalizations: []
};

const mapStateToProps = (state) => ({
    listTranslations: state.translations.listTranslations,
    listLocalizations: state.localizations.listLocalizations,
    errorMsg: state.translations.error,
});

const mapDispatchToProps = {
    listTranslationsRequest,
    createTranslationRequest,
    updateTranslationRequest,
    hideTranslationRequest,
    listLocalizationsRequest,
};

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Translations));