import styled from '@emotion/styled/macro';
import { useCreateUpstream } from 'Api';
import {
  SoloFormDropdown,
  SoloFormInput,
  SoloFormTypeahead
} from 'Components/Common/Form/SoloFormField';
import {
  SoloFormTemplate,
  InputRow,
  Footer
} from 'Components/Common/Form/SoloFormTemplate';
import { Field, Formik } from 'formik';
import { NamespacesContext } from 'GlooIApp';
import {
  UpstreamSpec as AwsUpstreamSpec,
  LambdaFunctionSpec
} from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/aws_pb';
import { UpstreamSpec as AzureUpstreamSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/azure/azure_pb';
import { UpstreamSpec as ConsulUpstreamSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/consul/consul_pb';
import { UpstreamSpec as KubeUpstreamSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/kubernetes/kubernetes_pb';
import { ServiceSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/service_spec_pb';
import { UpstreamSpec as StaticUpstreamSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/static/static_pb';
import { ResourceRef } from 'proto/github.com/solo-io/solo-kit/api/v1/ref_pb';
import {
  CreateUpstreamRequest,
  UpstreamInput
} from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/upstream_pb';
import * as React from 'react';
import { UPSTREAM_SPEC_TYPES, UPSTREAM_TYPES_CREATE } from 'utils/upstreamHelpers';
import * as yup from 'yup';
import { awsInitialValues, AwsUpstreamForm } from './AwsUpstreamForm';
import { azureInitialValues, AzureUpstreamForm } from './AzureUpstreamForm';
import {
  consulInitialValues,
  ConsulUpstreamForm,
  consulValidationSchema
} from './ConsulUpstreamForm';
import { kubeInitialValues, KubeUpstreamForm } from './KubeUpstreamForm';
import { staticInitialValues, StaticUpstreamForm } from './StaticUpstreamForm';
import { SoloButton } from 'Components/Common/SoloButton';
interface Props {}

const FormContainer = styled.div`
  display: flex;
  flex-direction: column;
`;

// TODO: better way to include all initial values?
export const initialValues = {
  name: '',
  type: '',
  namespace: 'gloo-system',
  ...awsInitialValues,
  ...kubeInitialValues,
  ...staticInitialValues,
  ...azureInitialValues,
  ...consulInitialValues
};

// TODO combine validation schemas
const validationSchema = yup.object().shape({
  name: yup
    .string()
    .required('Upstream name is required')
    .min(2, `Name can't be that short`),
  namespace: yup.string().required('Namespace is required'),
  type: yup.string().required('Must specify an upstream type'),
  awsRegion: yup.string(),
  awsSecretRef: yup.object().shape({
    name: yup.string().required('A secret is required'),
    namespace: yup.string()
  })
});

export const CreateUpstreamForm = (props: Props) => {
  const namespaces = React.useContext(NamespacesContext);

  const { refetch: makeRequest } = useCreateUpstream(null);

  // grpc request
  function createUpstream(values: typeof initialValues) {
    const newUpstreamReq = new CreateUpstreamRequest();
    const usInput = new UpstreamInput();

    const usResourceRef = new ResourceRef();

    usResourceRef.setName(values.name);
    usResourceRef.setNamespace(values.namespace);

    usInput.setRef(usResourceRef);

    //TODO: set up correct upstream spec
    // TODO: validation for specific fields
    switch (values.type) {
      case UPSTREAM_SPEC_TYPES.AWS:
        const awsSpec = new AwsUpstreamSpec();
        awsSpec.setRegion(values.awsRegion);
        const awsSecretRef = new ResourceRef();
        awsSecretRef.setName(values.awsSecretRef.name);
        awsSecretRef.setNamespace(values.awsSecretRef.namespace);
        awsSpec.setSecretRef(awsSecretRef);
        usInput.setAws(awsSpec);
        break;
      case UPSTREAM_SPEC_TYPES.AZURE:
        const azureSpec = new AzureUpstreamSpec();
        const azureSecretRef = new ResourceRef();
        azureSecretRef.setName(values.azureSecretRef.name);
        azureSecretRef.setNamespace(values.azureSecretRef.namespace);
        azureSpec.setSecretRef(azureSecretRef);
        azureSpec.setFunctionAppName(values.azureFunctionAppName);
        usInput.setAzure(azureSpec);
        break;
      case UPSTREAM_SPEC_TYPES.KUBE:
        const kubeSpec = new KubeUpstreamSpec();
        kubeSpec.setServiceName(values.kubeServiceName);
        kubeSpec.setServiceNamespace(values.kubeServiceNamespace);
        kubeSpec.setServicePort(values.kubeServicePort);
        usInput.setKube(kubeSpec);
        break;
      case UPSTREAM_SPEC_TYPES.STATIC:
        const staticSpec = new StaticUpstreamSpec();
        staticSpec.setUseTls(values.staticUseTls);
        usInput.setStatic(staticSpec);
        break;
      case UPSTREAM_SPEC_TYPES.CONSUL:
        const consulSpec = new ConsulUpstreamSpec();
        consulSpec.setServiceName(values.consulServiceName);
        consulSpec.setServiceTagsList(values.consulServiceTagsList);
        consulSpec.setConnectEnabled(values.consulConnectEnabled);
        consulSpec.setDataCentersList(values.consulDataCentersList);
        const consulServiceSpec = new ServiceSpec();
        consulSpec.setServiceSpec(consulServiceSpec);
        usInput.setConsul(consulSpec);
      default:
        break;
    }

    newUpstreamReq.setInput(usInput);
    makeRequest(newUpstreamReq);
  }

  return (
    <Formik
      initialValues={initialValues}
      validationSchema={validationSchema}
      onSubmit={createUpstream}>
      {formik => (
        <FormContainer>
          <SoloFormTemplate>
            <InputRow>
              <div>
                <SoloFormInput
                  name='name'
                  title='Upstream Name'
                  placeholder='Upstream Name'
                />
              </div>
              <div>
                <SoloFormDropdown
                  name='type'
                  title='Upstream Type'
                  placeholder='Type'
                  options={UPSTREAM_TYPES_CREATE}
                />
              </div>
              <div>
                <SoloFormTypeahead
                  name='namespace'
                  title='Upstream Namespace'
                  defaultValue='gloo-system'
                  presetOptions={namespaces}
                />
              </div>
            </InputRow>
          </SoloFormTemplate>
          {formik.values.type === UPSTREAM_SPEC_TYPES.AWS && (
            <AwsUpstreamForm />
          )}
          {formik.values.type === UPSTREAM_SPEC_TYPES.KUBE && (
            <KubeUpstreamForm />
          )}
          {formik.values.type === UPSTREAM_SPEC_TYPES.STATIC && (
            <StaticUpstreamForm />
          )}
          {formik.values.type === UPSTREAM_SPEC_TYPES.AZURE && (
            <AzureUpstreamForm />
          )}
          {formik.values.type === UPSTREAM_SPEC_TYPES.CONSUL && (
            <ConsulUpstreamForm />
          )}
          <Footer>
            <SoloButton
              onClick={e => formik.handleSubmit(e)}
              text='Create Upstream'
              disabled={formik.isSubmitting}
            />
          </Footer>
        </FormContainer>
      )}
    </Formik>
  );
};
