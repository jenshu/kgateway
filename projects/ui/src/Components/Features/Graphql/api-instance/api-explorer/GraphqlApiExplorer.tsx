import { QuestionCircleOutlined } from '@ant-design/icons';
import { Global } from '@emotion/core';
import styled from '@emotion/styled';
import { createGraphiQLFetcher } from '@graphiql/toolkit';
import { Tooltip } from 'antd';
import { useGetGraphqlApiDetails, useListVirtualServices } from 'API/hooks';
import { ReactComponent as WarningExclamation } from 'assets/big-warning-exclamation.svg';
import { ReactComponent as CopyIcon } from 'assets/document.svg';
import { SoloInput } from 'Components/Common/SoloInput';
import { Fetcher, GraphiQL } from 'graphiql';
import { buildSchema } from 'graphql';
import * as React from 'react';
import { ChangeEvent, useMemo, useRef, useState } from 'react';
import { useParams } from 'react-router';
import { colors } from 'Styles/colors';
import { copyTextToClipboard } from 'utils';
import {
  StatusHealth,
  WarningCircle,
} from '../../../Overview/OverviewBoxSummary';
import graphiqlCustomStyles from './GraphqlApiExplorer.style';

const Wrapper = styled.div`
  background: white;
`;

const StyledContainer = styled.div`
  height: 70vh;
`;

const GqlInputContainer = styled.div`
  padding: 15px 10px;
  border-bottom: 1px solid ${colors.marchGrey};
  display: flex;
`;

const GqlInputWrapper = styled.div`
  flex-basis: min-content;
  display: flex;
  flex-direction: row;
`;

const LabelTextWrapper = styled.div<{ hasError: boolean }>`
  width: 100%;
  label {
    width: 100%;
    margin-right: 10px;
    color: ${props => (props.hasError ? colors.sunGold : 'black')};
  }
  input {
    width: 250px;
  }
`;

const StyledQuestionMark = styled(QuestionCircleOutlined)`
  margin-top: 3px;
  display: inline-flex;
  &:hover {
    cursor: pointer;
  }
`;

const CodeWrapper = styled.div`
  code {
    &:hover {
      cursor: pointer;
      color: ${colors.aprilGrey};
      fill: ${colors.aprilGrey};
    }
  }
  p {
    padding: 10px 0;
  }
`;

const Copied = styled.span`
  display: inline-block;
  margin-left: 10px;
`;

const GQL_STORAGE_KEY = 'gqlStorageKey';

const StyledCopyIcon = styled(CopyIcon)`
  color: white;
  display: inline-block;
  margin-left: 10px;
  fill: white;
`;

const getGqlStorage = () => {
  return (
    localStorage.getItem(GQL_STORAGE_KEY) || 'http://localhost:8080/graphql'
  );
};

const setGqlStorage = (value: string) => {
  localStorage.setItem(GQL_STORAGE_KEY, value);
};

export const GraphqlApiExplorer = () => {
  const { graphqlApiName, graphqlApiNamespace, graphqlApiClusterName } =
    useParams();
  const [gqlError, setGqlError] = useState('');
  const [refetch, setRefetch] = useState(false);
  const [url, setUrl] = useState(getGqlStorage());
  const [showTooltip, setShowTooltip] = useState(false);
  const [copyingKubectl, setCopyingKubectl] = useState(false);
  const [copyingProxy, setCopyingProxy] = useState(false);
  const [showUrlBar, setShowUrlBar] = useState(false);

  const {
    data: graphqlApi,
    error: graphqlApiError,
    mutate,
  } = useGetGraphqlApiDetails({
    name: graphqlApiName,
    namespace: graphqlApiNamespace,
    clusterName: graphqlApiClusterName,
  });

  const changeUrl = (value: string) => {
    setGqlStorage(value);
    setUrl(value);
  };

  const copyKubectlCommand = async () => {
    setCopyingKubectl(true);
    const text =
      'kubectl port-forward -n gloo-system deploy/gateway-proxy 8080';
    copyTextToClipboard(text).finally(() => {
      setTimeout(() => {
        setCopyingKubectl(false);
      }, 2000);
    });
  };

  const copyGlooctlCommand = async () => {
    setCopyingProxy(true);
    const text = 'glooctl proxy url';
    copyTextToClipboard(text).finally(() => {
      setTimeout(() => {
        setCopyingProxy(false);
      }, 2000);
    });
  };

  // If we need the custom fetcher, we can add `schemaFetcher` to the document.
  let gqlFetcher: Fetcher = createGraphiQLFetcher({
    url,
    schemaFetcher: async graphQLParams => {
      try {
        setRefetch(false);
        setGqlError('');
        const data = await fetch(url, {
          method: 'POST',
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(graphQLParams),
          credentials: 'same-origin',
        });
        return data.json().catch(() => data.text());
      } catch (error: any) {
        setGqlError(error.message);
      }
    },
  });

  const graphiqlRef = useRef<null | GraphiQL>(null);

  const { data: virtualServices, error: virtualServicesError } =
    useListVirtualServices();

  const correspondingVirtualServices = useMemo(
    () =>
      virtualServices?.filter(vs =>
        vs.spec?.virtualHost?.routesList.some(
          route =>
            route?.graphqlApiRef?.name === graphqlApiName &&
            route?.graphqlApiRef?.namespace === graphqlApiNamespace
        )
      ),
    [virtualServices, graphqlApiName, graphqlApiNamespace]
  );

  let executableSchema;

  if (graphqlApi?.spec?.executableSchema?.schemaDefinition) {
    const schemaDef = graphqlApi.spec.executableSchema.schemaDefinition;
    executableSchema = buildSchema(schemaDef, {
      assumeValidSDL: true,
    });
  }

  const handlePrettifyQuery = () => {
    graphiqlRef?.current?.handlePrettifyQuery();
  };
  const changeHost = (e: ChangeEvent<HTMLInputElement>) => {
    setRefetch(true);
    changeUrl(e.currentTarget.value);
  };

  const toggleUrlBar = () => {
    setShowUrlBar(!showUrlBar);
  };

  // The operation name === the tab name.
  // This comes from the actual GraphQL operation.
  //   e.g. query Test {} will have the tab name "Test".
  //   The example has "query Example { }" in it.
  const [opName, setOpName] = useState('Example');

  // TODO:  We can hide and show elements based on what we get back.
  //        The schema will only refetch if the executable schema is undefined.
  if (correspondingVirtualServices === undefined) return null;
  return correspondingVirtualServices.length > 0 ? (
    <Wrapper>
      <Global styles={graphiqlCustomStyles} />
      {Boolean(gqlError) || showUrlBar ? (
        <GqlInputContainer>
          <GqlInputWrapper>
            <LabelTextWrapper hasError={Boolean(gqlError)}>
              <SoloInput
                title={`${
                  Boolean(gqlError)
                    ? 'Failed to fetch Graphql service.  Update the host to attempt again.'
                    : 'Current url'
                }`}
                value={url}
                onChange={changeHost}
              />
            </LabelTextWrapper>
            <Tooltip
              title={
                <CodeWrapper>
                  <p>
                    Endpoint URL for the gateway proxy. The default URL can be
                    used if you port forward with the following command:
                  </p>
                  <p
                    className='copy'
                    title='copy command'
                    onClick={copyKubectlCommand}>
                    <code>
                      <i>
                        kubectl port-forward -n gloo-system deploy/gateway-proxy
                        8080
                      </i>
                      {copyingKubectl ? (
                        <Copied>copied!</Copied>
                      ) : (
                        <StyledCopyIcon />
                      )}
                    </code>
                  </p>
                  <p>
                    Depending on your installation, you can also use the
                    following glooctl command:
                  </p>
                  <p
                    className='copy'
                    title='copy command'
                    onClick={copyGlooctlCommand}>
                    <code>
                      <i>glooctl proxy url</i>
                      {copyingProxy ? (
                        <Copied>copied!</Copied>
                      ) : (
                        <StyledCopyIcon />
                      )}
                    </code>
                  </p>
                </CodeWrapper>
              }
              trigger='hover'
              visible={showTooltip}
              onVisibleChange={() => {
                setShowTooltip(!showTooltip);
              }}>
              <StyledQuestionMark />
            </Tooltip>
          </GqlInputWrapper>
        </GqlInputContainer>
      ) : null}
      <StyledContainer>
        <GraphiQL
          ref={graphiqlRef}
          defaultQuery={`query Example {

}

# Welcome to GraphiQL, an in-browser tool for
# writing, validating, and testing GraphQL queries.
#
# Type queries into this side of the screen, and you
# will see intelligent typeaheads aware of the current
# GraphQL type schema and live syntax and
# validation errors highlighted within the text.
#
# GraphQL queries typically start with a "{" character.
# Lines that start with a # are ignored.
# The name of the query on the first line of each tab
# is the title of that tab.
#
# An example GraphQL query might look like:
#     query Example {
#       field(arg: "value") {
#         subField
#       }
#     }
#
# Keyboard shortcuts:
#     Prettify Query:    Shift-Ctrl-P
#     Merge Query:     Shift-Ctrl-M
#     Run Query:        Ctrl-Enter
#     Auto Complete:  Ctrl-Space

`}
          variables={'{}'}
          tabs={true}
          operationName={opName}
          onEditOperationName={s => setOpName(s)}
          schema={!refetch ? executableSchema : undefined}
          fetcher={gqlFetcher}>
          <GraphiQL.Toolbar>
            <GraphiQL.Button
              onClick={handlePrettifyQuery}
              label='Prettify'
              title='Prettify Query (Shift-Ctrl-P)'
            />
            <GraphiQL.Button
              onClick={toggleUrlBar}
              label={showUrlBar ? 'Hide Url Bar' : 'Show Url Bar'}
              title='Show/Hide Url Bar'
            />
          </GraphiQL.Toolbar>
        </GraphiQL>
      </StyledContainer>
    </Wrapper>
  ) : (
    <div className='overflow-hidden bg-white rounded-lg shadow'>
      <div className='px-4 py-5 sm:p-6'>
        <StatusHealth isWarning>
          <div>
            <WarningCircle>
              <WarningExclamation />
            </WarningCircle>
          </div>
          <div>
            <>
              <div className='text-xl '>Explorer unavailable</div>
              <div className='text-lg '>
                There is no Virtual Service that exposes this GraphQL endpoint
              </div>
            </>
          </div>
        </StatusHealth>
      </div>
    </div>
  );
};
