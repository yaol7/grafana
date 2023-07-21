import React from 'react';
import { RouteChildrenProps } from 'react-router-dom';

import { EntityNotFound } from 'app/core/components/PageNotFound/EntityNotFound';

import { useAlertmanagerConfig } from '../../hooks/useAlertmanagerConfig';
import { useAlertmanager } from '../../state/AlertmanagerContext';
import { EditTemplateView } from '../receivers/EditTemplateView';

type Props = RouteChildrenProps<{ name: string }>;

const EditMessageTemplate = ({ match }: Props) => {
  const { selectedAlertmanager } = useAlertmanager();
  const { data, isLoading, error } = useAlertmanagerConfig(selectedAlertmanager);

  const name = match?.params.name;
  if (!name) {
    return <EntityNotFound entity="Message template" />;
  }

  if (isLoading && !data) {
    return 'loading...';
  }

  // TODO decent error handling
  if (error) {
    return String(error);
  }

  if (!data) {
    return null;
  }

  return (
    <EditTemplateView
      alertManagerSourceName={selectedAlertmanager!}
      config={data}
      templateName={decodeURIComponent(name)}
    />
  );
};

export default EditMessageTemplate;