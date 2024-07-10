import { Box } from '@grafana/ui';
import { t, Trans } from 'app/core/internationalization';

import { InfoItem } from '../../shared/InfoItem';
import { MigrationTokenPane } from '../MigrationTokenPane/MigrationTokenPane';

export const InfoPane = () => {
  return (
    <Box alignItems="flex-start" display="flex" direction="column" gap={2}>
      <InfoItem title={t('migrate-to-cloud.migrate-to-this-stack.title', 'Let us help you migrate to this stack')}>
        <Trans i18nKey="migrate-to-cloud.migrate-to-this-stack.body">
          You can migrate some resources from your self-managed Grafana installation to this cloud stack. To do this
          securely, you&apos;ll need to generate a migration token. Your self-managed instance will use the token to
          authenticate with this cloud stack.
        </Trans>
      </InfoItem>
      <MigrationTokenPane />
    </Box>
  );
};
