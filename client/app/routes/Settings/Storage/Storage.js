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
import {connect} from 'react-redux';
import {updateSettingStorageRequest} from "../../../redux/settings/actions";

export class Storage extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            storage: this.props.storage,
        };
    }

    handleChange = (e) => {
        const {name, value} = e.target;
        const s = {...this.state.storage, [name]: value};
        this.setState(function (previousState) {
            return {...previousState, storage: s};
        });
    };

    handleUpdateStorage = (e) => {
        e.preventDefault();
        const {use} = this.state.storage;
        this.props.updateSettingStorageRequest(parseInt(use));
    };

    render() {
        const {use} = this.state.storage;
        return (
            <React.Fragment>
                <Container>
                    <HeaderMain
                        title="Storage Edit"
                        className="mb-5 mt-4"
                    />
                    <Row>
                        <Col lg={12}>
                            <Card className="mb-3">
                                <CardBody>
                                    <div className="d-flex mb-4">
                                        <CardTitle tag="h6">
                                            Storage Edit
                                        </CardTitle>
                                        <span className="ml-auto align-self-start small">
                                    Fields mark as <span className="text-danger">*</span> is required.
                                </span>
                                    </div>
                                    <Form>
                                        <FormGroup row>
                                            <Label for="use" sm={3} className="text-right">
                                                <span className="text-danger">*</span> Storage
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="select"
                                                    name="use"
                                                    value={use}
                                                    onChange={(e) => this.handleChange(e)}
                                                >
                                                    <option value={0}>Local Storage</option>
                                                    <option value={1}>S3 Storage</option>
                                                </Input>
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label sm={3}>
                                            </Label>
                                        </FormGroup>
                                    </Form>
                                </CardBody>
                                <CardFooter className="small text-right">
                                    <Button onClick={(e) => this.handleUpdateStorage(e)} color="primary">
                                        Update Storage
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

Storage.propTypes = {
    updateSettingStorageRequest: PropTypes.func.isRequired,
    storage: PropTypes.object
};

const mapStateToProps = (state) => ({
    storage: state.settings.settings.storage,
});

const mapDispatchToProps = {
    updateSettingStorageRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(Storage);