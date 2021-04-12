import React from 'react';

import {
    Container,
    Row,
    Col,
} from './../../components';

import {HeaderMain} from '../components/HeaderMain';
import {withRouter} from 'react-router-dom';
import * as PropTypes from 'prop-types';
import {Table} from './Table';
import {connect} from 'react-redux';
import {userListRequest} from '../../redux/users/actions';

class Users extends React.Component {
    constructor(props) {
        super(props);

        this.props.userListRequest();
    }

    render() {
        const {listUsers} = this.props;
        return (
            <Container>
                <HeaderMain
                    title="Users"
                    className="mb-5 mt-4"
                />
                <Row>
                    <Col lg={12}>
                        <Table
                            history={this.props.history}
                            listUsers={listUsers}
                        />
                    </Col>
                </Row>
            </Container>
        );
    }
}

Users.propTypes = {
    listUsers: PropTypes.array,
    history: PropTypes.shape({
        push: PropTypes.func.isRequired,
    }).isRequired,
    userListRequest: PropTypes.func.isRequired,
};

Users.defaultProps = {
    listUsers: []
};

const mapStateToProps = (state) => ({
    listUsers: state.users.users,
});

const mapDispatchToProps = {
    userListRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(withRouter(Users));
