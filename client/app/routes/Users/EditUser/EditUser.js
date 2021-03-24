import React from 'react';

import {
    Container,
    Row,
    Card,
    Button,
    CardBody,
    Form, FormGroup, Label, Input, Col, CustomInput, CardFooter
} from '../../../components';

import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import {HeaderMain} from '../../components/HeaderMain';
import {updateUserProfileRequest} from '../../../redux/profile/actions';

export class EditUser extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            user: this.props.location.state
        };
    }

    handleChange = (e) => {
        const {name, value} = e.target;
        const o = {...this.state.user, [name]: value};
        this.setState({user: o});
        console.log(this.state.user);
    };

    handleUpdateUser = () => {
        const {id, first_name, second_name, email, role, status} = this.state.user;
        this.props.updateUserProfileRequest(first_name, second_name, email, role, status, id);
        this.props.history.goBack();
    };

    render() {
        const {first_name, second_name, email, role, status} = this.state.user;

        return (
            <React.Fragment>
                <Container>
                    <HeaderMain
                        title="User Edit"
                        className="mb-5 mt-4"
                    />
                    <Row>
                        <Col lg={12}>
                            <Card>
                                <CardBody>
                                    <Form>
                                        <div className="small mt-4 mb-3">
                                            Personal Information
                                        </div>
                                        <FormGroup row>
                                            <Label for="firstName" sm={3} className="text-right">
                                                First Name
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="text"
                                                    name="first_name"
                                                    placeholder="First Name..."
                                                    value={first_name}
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label for="lastName" sm={3} className="text-right">
                                                Second Name
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="text"
                                                    name="second_name"
                                                    placeholder="Last Name..."
                                                    value={second_name}
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label for="lastName" sm={3} className="text-right">
                                                User Email
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="text"
                                                    name="email"
                                                    placeholder="Email..."
                                                    value={email}
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <div className="small mt-5 mb-3">
                                            Availability
                                        </div>
                                        <FormGroup row>
                                            <Label for="jobType" sm={3} className="text-right">
                                                User Role
                                            </Label>
                                            <Col sm={8}>
                                                <CustomInput
                                                    type="select"
                                                    id="role"
                                                    name="role"
                                                    value={role}
                                                    onChange={(e) => this.handleChange(e)}
                                                >
                                                    <option value="">Select...</option>
                                                    <option value="admin">Admin</option>
                                                    <option value="translator">Translator</option>
                                                    <option value="moderator">Moderator</option>
                                                </CustomInput>
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label for="jobType" sm={3} className="text-right">
                                                User Status
                                            </Label>
                                            <Col sm={8}>
                                                <CustomInput
                                                    type="select"
                                                    id="status"
                                                    name="status"
                                                    value={status}
                                                    onChange={(e) => this.handleChange(e)}
                                                >
                                                    <option value="">Select...</option>
                                                    <option value="active">Active</option>
                                                    <option value="not active">Not Active</option>
                                                    <option value="banned">Banned</option>
                                                </CustomInput>
                                            </Col>
                                        </FormGroup>
                                    </Form>
                                </CardBody>
                                <CardFooter className="text-right">
                                    <Button onClick={this.handleUpdateUser} color="primary">
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

EditUser.propTypes = {
    history: PropTypes.object.isRequired,
    location: PropTypes.object.isRequired,
    updateUserProfileRequest: PropTypes.func.isRequired,
};

export default connect(null, {updateUserProfileRequest})(EditUser);