import React from 'react';

import {
    Container,
    Row,
    Col,
} from './../../components';

import {HeaderMain} from '../components/HeaderMain';

import {CategoryTable} from './Table';
import {Action} from './Action';
import {Media} from 'reactstrap';
import {toast, ToastContainer} from 'react-toastify';
import {connect} from 'react-redux';
import * as PropTypes from 'prop-types';
import {
    createCategoryRequest,
    deleteCategoryRequest,
    listCategoriesRequest,
    updateCategoryRequest
} from '../../redux/categories/actions';

const defaultChosenCategory = {id: 0, name: ''};

class Categories extends React.Component {
    constructor(props) {
        super(props);

        this.props.listCategoriesRequest();
        this.state = {
            chosenCategory: defaultChosenCategory,
        };
    }

    handleChange = (e) => {
        const {name, value} = e.target;
        const o = {...this.state.chosenCategory, [name]: value};
        this.setState(function (previousState) {
            return {
                ...previousState,
                chosenCategory: o
            };
        });
    };

    handleChosenCategory = (index) => {
        this.setState({
            ...this.state,
            chosenCategory: this.props.listCategories[index]
        });
    };

    handleResetChosenCategory = () => {
        this.setState({
            ...this.state,
            chosenCategory: defaultChosenCategory
        });
    };

    handleDeleteCategory = (id) => {
        this.props.deleteCategoryRequest(id);
    };

    handleSubmitCategory = (e) => {
        e.preventDefault();
        const {id, name} = this.state.chosenCategory;
        if (name.trim() === '') {
            toast.success(contentSuccess('Value not valid!'));
            return;
        }
        if (id > 0) {
            this.props.updateCategoryRequest(id, name);
        } else {
            this.props.createCategoryRequest(name);
        }
        this.handleResetChosenCategory();
        // toast.success(contentSuccess('Localization successful added!'));
    };

    render() {
        const {listCategories} = this.props;
        return (
            <Container>
                <HeaderMain
                    title="Categories"
                    className="mb-5 mt-4"
                />
                <Row>
                    <Col lg={8}>
                        <CategoryTable
                            listCategories={listCategories}
                            handleChosenCategory={this.handleChosenCategory}
                            handleDeleteCategory={this.handleDeleteCategory}
                        />
                    </Col>
                    <Col lg={4}>
                        <Action
                            chosenCategory={this.state.chosenCategory}
                            handleChange={this.handleChange}
                            handleResetChosenCategory={this.handleResetChosenCategory}
                            handleSubmitCategory={this.handleSubmitCategory}
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

Categories.propTypes = {
    listCategories: PropTypes.array,
    listCategoriesRequest: PropTypes.func.isRequired,
    createCategoryRequest: PropTypes.func.isRequired,
    updateCategoryRequest: PropTypes.func.isRequired,
    deleteCategoryRequest: PropTypes.func.isRequired,
    errorMsg: PropTypes.string,
};

const mapStateToProps = (state) => ({
    listCategories: state.categories.listCategories,
    errorMsg: state.categories.error,
});

const mapDispatchToProps = {
    listCategoriesRequest,
    createCategoryRequest,
    updateCategoryRequest,
    deleteCategoryRequest
};

export default connect(mapStateToProps, mapDispatchToProps)(Categories);