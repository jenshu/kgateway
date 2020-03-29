import SwaggerUI from 'swagger-ui-react';
import 'swagger-ui-react/swagger-ui.css';
import { css } from '@emotion/core';
import React from 'react';
import { ReactComponent as HealthScoreIcon } from 'assets/health-score-icon.svg';
import { ReactComponent as DevPortalIcon } from 'assets/developer-portal-icon.svg';
import { ReactComponent as CodeIcon } from 'assets/code-icon.svg';
import { ReactComponent as UserIcon } from 'assets/user-icon.svg';

import { SectionCard } from 'Components/Common/SectionCard';
import { StatusTile } from 'Components/Common/DisplayOnly/StatusTile';
import { ReactComponent as PortalIcon } from 'assets/portal-icon.svg';
import { Container, Header, HealthScoreContainer } from '../Admin/AdminLanding';
import { healthConstants } from 'Styles';
import { TallyInformationDisplay } from 'Components/Common/DisplayOnly/TallyInformationDisplay';
import { SwaggerExplorer } from './SwaggerExplorer';
import { ErrorBoundary } from '../Errors/ErrorBoundary';
import useSWR from 'swr';
import { portalApi, apiDocApi, userApi, groupApi } from './api';
import { formatHealthStatus } from './portals/PortalsListing';
import { Status } from 'proto/solo-kit/api/v1/status_pb';

export const DevPortalOverview = () => {
  const { data: portalsList, error: portalsListError } = useSWR(
    'listPortals',
    portalApi.listPortals
  );

  const { data: apiDocsList, error: apiDocsError } = useSWR(
    'listApiDocs',
    apiDocApi.listApiDocs
  );

  const { data: userList, error: userError } = useSWR(
    'listUsers',
    userApi.listUsers
  );

  const { data: groupList, error: groupError } = useSWR(
    'listGroups',
    groupApi.listGroups
  );

  if (!portalsList || !apiDocsList || !userList || !groupList) {
    return <div>Loading...</div>;
  }

  let publishedApiDocsCount = 0;
  portalsList.forEach(portal => {
    publishedApiDocsCount +=
      portal.spec?.publishApiDocs?.matchLabelsMap.length || 0;
  });

  let endpointCount = 0;
  apiDocsList.forEach(apiDoc => {
    endpointCount += apiDoc.status?.numberOfEndpoints || 0;
  });

  let portalErrorPresent = portalsList.some(
    portal =>
      formatHealthStatus(portal.status?.state) === Status.State.PENDING ||
      formatHealthStatus(portal.status?.state) === Status.State.REJECTED
  );

  let apiDocErrorPresent = apiDocsList.some(
    apiDoc =>
      formatHealthStatus(apiDoc.status?.state) === Status.State.PENDING ||
      formatHealthStatus(apiDoc.status?.state) === Status.State.REJECTED
  );

  return (
    <ErrorBoundary
      fallback={<div>There was an error with the Dev Portal section</div>}>
      <Container>
        <Header>
          <div>
            <div className='text-2xl '>Developer Portal</div>
            <div className='text-lg text-gray-700'>
              Create and share APIs through custom-branded client portals.{' '}
            </div>
          </div>
          <HealthScoreContainer health={healthConstants.Good.value}>
            <span className='text-blue-500'>
              <DevPortalIcon className='w-32 h-32 fill-current' />
            </span>
          </HealthScoreContainer>
        </Header>
        <div className='grid grid-cols-3 px-4 py-5 sm:p-6'>
          <StatusTile
            exploreMoreLink={{
              prompt: 'View Portals',
              link: '/dev-portal/portals'
            }}
            healthStatus={
              portalErrorPresent
                ? healthConstants.Pending.value
                : healthConstants.Good.value
            }
            titleText='Portals'
            description='Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.'
            titleIcon={
              <span className='text-blue-500'>
                <PortalIcon className='w-8 h-8 fill-current ' />
              </span>
            }>
            {portalsList.length > 0 ? (
              <div className='grid grid-cols-2 gap-4 '>
                <TallyInformationDisplay
                  tallyCount={portalsList.length}
                  tallyDescription={`portals`}
                  color='blue'
                />
                <TallyInformationDisplay
                  tallyCount={publishedApiDocsCount}
                  tallyDescription={`published APIs`}
                  color='blue'
                />
              </div>
            ) : (
              <TallyInformationDisplay
                tallyDescription={`You have no Portals detected. Get started by creating a Portal!`}
                color='blue'
              />
            )}
          </StatusTile>
          <StatusTile
            exploreMoreLink={{
              prompt: 'View APIs',
              link: '/dev-portal/apis'
            }}
            healthStatus={
              apiDocErrorPresent
                ? healthConstants.Pending.value
                : healthConstants.Good.value
            }
            titleText='APIs'
            description='Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.'
            titleIcon={
              <span className='text-blue-500'>
                <CodeIcon className='w-8 h-8 fill-current' />
              </span>
            }>
            {apiDocsList.length > 0 ? (
              <div className='grid grid-cols-2 gap-4 '>
                <TallyInformationDisplay
                  tallyCount={apiDocsList.length}
                  tallyDescription={`APIs`}
                  color='blue'
                />
                <TallyInformationDisplay
                  tallyCount={endpointCount}
                  tallyDescription={`Endpoints`}
                  color='blue'
                />
              </div>
            ) : (
              <TallyInformationDisplay
                tallyDescription={`You have no Portals detected. Get started by creating a Portal!`}
                color='blue'
              />
            )}
          </StatusTile>

          <StatusTile
            exploreMoreLink={{
              prompt: 'View Users & Groups',
              link: '/dev-portal/users'
            }}
            healthStatus={1}
            titleText='Users & Groups'
            description='Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.'
            titleIcon={
              <span className='text-blue-500'>
                <UserIcon className='w-8 h-8 fill-current ' />
              </span>
            }>
            {userList.length > 0 ? (
              <div className='grid grid-cols-2 gap-4 '>
                <TallyInformationDisplay
                  tallyCount={userList.length}
                  tallyDescription='Users'
                  color='blue'
                />
                <TallyInformationDisplay
                  tallyCount={groupList.length}
                  tallyDescription='Groups'
                  color='blue'
                />
              </div>
            ) : (
              <TallyInformationDisplay
                tallyDescription={`You have no users or groups detected. Get started by creating a user or group!`}
                color='blue'
              />
            )}
          </StatusTile>
        </div>
        <div>
          <div className='text-2xl text-gray-900 '>
            Developer Portal Documentation
          </div>
        </div>
        <div className='grid grid-cols-3 px-4 py-5 sm:p-6'>
          <StatusTile
            exploreMoreLink={{ prompt: 'View Documentation', link: '' }}
            titleText='Creating a Portal'
            description='Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.'></StatusTile>
          <StatusTile
            exploreMoreLink={{ prompt: 'View Documentation', link: '' }}
            titleText='Generating a Gloo API'
            description='Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.'></StatusTile>

          <StatusTile
            exploreMoreLink={{ prompt: 'View Documentation', link: '' }}
            titleText='Granting User Access'
            description='Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.'></StatusTile>
        </div>
      </Container>
    </ErrorBoundary>
  );
};
