import {
    Button,
    Card,
    CardBody,
    CardTitle,
    Form,
    FormGroup,
    Input,
} from '../../components';
import React from 'react';
import PropTypes from 'prop-types';
import {Typeahead} from 'react-bootstrap-typeahead';

export class Action extends React.Component {
    render() {
        const isEdit = this.props.chosenLocalization.id > 0;
        const {
            listLanguages,
            chosenLocalization,
            handleChooseLocalization,
            handleSubmitLocalization,
        } = this.props;

        return (
            <Card className="mb-3">
                <CardBody>
                    <CardTitle tag="h6" className="mb-4">
                        Create Localization
                    </CardTitle>
                    <Form>
                        <FormGroup>
                            <Typeahead
                                clearButton
                                id="list-available-localizations-for-id"
                                defaultSelected={[]}
                                labelKey="name"
                                options={listLanguages}
                                placeholder="Choose a localization..."
                                onChange={handleChooseLocalization}
                            />
                        </FormGroup>
                        <FormGroup>
                            <Input
                                type="text"
                                value={chosenLocalization.code}
                                placeholder="Language code (ex: Czech)"
                                disabled={true}
                            />
                        </FormGroup>
                        <Button onClick={(e) => handleSubmitLocalization(e)} color="primary">
                            {
                                isEdit ? (`Edit`) : (`Create`)
                            }
                        </Button>
                    </Form>
                </CardBody>
            </Card>
        );
    }
}

Action.propTypes = {
    chosenLocalization: PropTypes.object.isRequired,
    handleChooseLocalization: PropTypes.func.isRequired,
    handleSubmitLocalization: PropTypes.func.isRequired,
    listLanguages: PropTypes.array.isRequired,
};
