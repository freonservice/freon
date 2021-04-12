import React from 'react';

import {
    Container,
    Row,
    Col,
} from './../../components';
import {ToastContainer} from 'react-toastify';
import * as PropTypes from 'prop-types';

import {HeaderMain} from '../components/HeaderMain';

import {connect} from 'react-redux';
import {IdentifierTable} from './Table';
import {Action} from './Action';
import {
    listIdentifiersRequest,
    createIdentifierRequest,
    updateIdentifierRequest,
    deleteIdentifierRequest
} from '../../redux/identifiers/actions';
import {listCategoriesRequest} from '../../redux/categories/actions';

const defaultChosenIdentifier = {id: 0, name: '', description: '', example_text: '', platforms: []};

class Identifiers extends React.Component {
    constructor(props) {
        super(props);

        this.props.listIdentifiersRequest();
        this.props.listCategoriesRequest();

        this.state = {
            chosenIdentifier: defaultChosenIdentifier,
        };
    }

    handleChangeIdentifierInformation = (e) => {
        const {name, value} = e.target;
        const o = {...this.state.chosenIdentifier, [name]: value};
        this.setState({chosenIdentifier: o});
    };

    handleChosenIdentifier = (id) => {
        this.setState(function () {
            return {chosenIdentifier: this.props.listIdentifiers.find(o => o.id === id) || {...defaultChosenIdentifier}};
        });
    };

    handleResetChosenIdentifier = () => {
        this.setState({chosenIdentifier: defaultChosenIdentifier});
    };

    render() {
        const {
            listIdentifiers, listCategories,
            deleteIdentifierRequest, updateIdentifierRequest, createIdentifierRequest
        } = this.props;

        const {chosenIdentifier} = this.state;

        return (
            <Container>
                <HeaderMain
                    title="Identifiers"
                    className="mb-5 mt-4"
                />
                <Row>
                    <Col lg={8}>
                        <IdentifierTable
                            listIdentifiers={listIdentifiers}
                            handleChosenIdentifier={this.handleChosenIdentifier}
                            deleteIdentifierRequest={deleteIdentifierRequest}
                        />
                    </Col>
                    <Col lg={4}>
                        <Action
                            leistCategories={listCategories}
                            handleChangeIdentifierInformation={this.handleChangeIdentifierInformation}
                            chosenIdentifier={chosenIdentifier}
                            updateIdentifierRequest={updateIdentifierRequest}
                            createIdentifierRequest={createIdentifierRequest}
                            handleResetChosenIdentifier={this.handleResetChosenIdentifier}
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

Identifiers.propTypes = {
    listIdentifiers: PropTypes.array,
    listCategories: PropTypes.array,
    listIdentifiersRequest: PropTypes.func.isRequired,
    createIdentifierRequest: PropTypes.func.isRequired,
    updateIdentifierRequest: PropTypes.func.isRequired,
    listCategoriesRequest: PropTypes.func.isRequired,
    deleteIdentifierRequest: PropTypes.func.isRequired,
    errorMsg: PropTypes.string,
};

Identifiers.defaultProps = {
    listIdentifiers: []
};

const mapStateToProps = (state) => ({
    listIdentifiers: state.identifiers.listIdentifiers,
    errorMsg: state.identifiers.error,
    listCategories: state.categories.listCategories,
});

const mapDispatchToProps = {
    listIdentifiersRequest,
    createIdentifierRequest,
    updateIdentifierRequest,
    listCategoriesRequest,
    deleteIdentifierRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(Identifiers);