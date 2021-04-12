import React from 'react';

import {
    Container,
    Row,
    Col,
    Input,
    Card,
    Button,
    CardBody,
    CardFooter,
    CardTitle,
    Label,
    FormGroup,
    Form
} from '../../../components';

import {HeaderMain} from '../../components/HeaderMain';
import * as PropTypes from 'prop-types';
import {updateUserPasswordRequest} from '../../../redux/profile/actions';
import {connect} from 'react-redux';

export class Security extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            changePassword: {old_password: '', new_password: '', repeat_password: ''}
        };
    }

    handleChange = (e) => {
        const {name, value} = e.target;
        const o = {...this.state.changePassword, [name]: value};
        this.setState(function (previousState) {
            return {...previousState, changePassword: o};
        });
    };

    handleUpdatePassword = (e) => {
        e.preventDefault();
        const {old_password, new_password, repeat_password} = this.state.changePassword;
        if (old_password.trim() === '' || new_password.trim() === '' || repeat_password.trim() === '' || new_password.length !== repeat_password.length) {
            alert('Password error');
            return;
        }
        this.props.updateUserPasswordRequest(old_password, new_password, repeat_password);
    };

    render() {
        return (
            <React.Fragment>
                <Container>
                    <HeaderMain
                        title="Security Edit"
                        className="mb-5 mt-4"
                    />
                    <Row>
                        <Col lg={12}>
                            <Card className="mb-3">
                                <CardBody>
                                    <div className="d-flex mb-4">
                                        <CardTitle tag="h6">
                                            Security Edit
                                        </CardTitle>
                                        <span className="ml-auto align-self-start small">
                                    Fields mark as <span className="text-danger">*</span> is required.
                                </span>
                                    </div>
                                    <Form>
                                        <FormGroup row>
                                            <Label for="oldPassword" sm={3} className="text-right">
                                                <span className="text-danger">*</span> Old Password
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="password"
                                                    name="old_password"
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label for="newPassword" sm={3} className="text-right">
                                                <span className="text-danger">*</span> New Password
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="password"
                                                    name="new_password"
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row className="mb-0">
                                            <Label for="confirmNewPassword" sm={3} className="text-right">
                                                <span className="text-danger">*</span> Confirm New Password
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="password"
                                                    name="repeat_password"
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label sm={3}>
                                            </Label>
                                        </FormGroup>
                                    </Form>
                                </CardBody>
                                <CardFooter className="small text-right">
                                    <Button onClick={(e) => this.handleUpdatePassword(e)} color="primary">
                                        Update Profile
                                    </Button>
                                </CardFooter>
                            </Card>
                        </Col>
                    </Row>
                </Container>
            </React.Fragment>
        );
    }
}

Security.propTypes = {
    updateUserPasswordRequest: PropTypes.func.isRequired,
};

const mapStateToProps = (state) => ({
    profile: state.profile.profile,
});

const mapDispatchToProps = {
    updateUserPasswordRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(Security);