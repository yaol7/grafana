import { rest } from 'msw';
import { setupServer, SetupServer } from 'msw/node';
import 'whatwg-fetch';

import { DataSourceInstanceSettings } from '@grafana/data';
import { setBackendSrv } from '@grafana/runtime';
import {
  PromBuildInfoResponse,
  PromRulesResponse,
  RulerRuleGroupDTO,
  RulerRulesConfigDTO,
} from 'app/types/unified-alerting-dto';

import { backendSrv } from '../../../core/services/backend_srv';
import {
  AlertmanagerConfig,
  AlertManagerCortexConfig,
  AlertmanagerReceiver,
  EmailConfig,
  MatcherOperator,
  Route,
} from '../../../plugins/datasource/alertmanager/types';

import { AlertingQueryResponse } from './state/AlertingQueryRunner';

class AlertmanagerConfigBuilder {
  private alertmanagerConfig: AlertmanagerConfig = { receivers: [] };

  addReceivers(configure: (builder: AlertmanagerReceiverBuilder) => void): AlertmanagerConfigBuilder {
    const receiverBuilder = new AlertmanagerReceiverBuilder();
    configure(receiverBuilder);
    this.alertmanagerConfig.receivers?.push(receiverBuilder.build());
    return this;
  }

  withRoute(configure: (routeBuilder: AlertmanagerRouteBuilder) => void): AlertmanagerConfigBuilder {
    const routeBuilder = new AlertmanagerRouteBuilder();
    configure(routeBuilder);

    this.alertmanagerConfig.route = routeBuilder.build();

    return this;
  }

  build() {
    return this.alertmanagerConfig;
  }
}

class AlertmanagerRouteBuilder {
  private route: Route = { routes: [], object_matchers: [] };

  withReceiver(receiver: string): AlertmanagerRouteBuilder {
    this.route.receiver = receiver;
    return this;
  }
  withoutReceiver(): AlertmanagerRouteBuilder {
    return this;
  }
  withEmptyReceiver(): AlertmanagerRouteBuilder {
    this.route.receiver = '';
    return this;
  }

  addRoute(configure: (builder: AlertmanagerRouteBuilder) => void): AlertmanagerRouteBuilder {
    const routeBuilder = new AlertmanagerRouteBuilder();
    configure(routeBuilder);
    this.route.routes?.push(routeBuilder.build());
    return this;
  }

  addMatcher(key: string, operator: MatcherOperator, value: string): AlertmanagerRouteBuilder {
    this.route.object_matchers?.push([key, operator, value]);
    return this;
  }

  build() {
    return this.route;
  }
}

class EmailConfigBuilder {
  private emailConfig: EmailConfig = { to: '' };

  withTo(to: string): EmailConfigBuilder {
    this.emailConfig.to = to;
    return this;
  }

  build() {
    return this.emailConfig;
  }
}

class AlertmanagerReceiverBuilder {
  private receiver: AlertmanagerReceiver = { name: '', email_configs: [] };

  withName(name: string): AlertmanagerReceiverBuilder {
    this.receiver.name = name;
    return this;
  }

  addEmailConfig(configure: (builder: EmailConfigBuilder) => void): AlertmanagerReceiverBuilder {
    const builder = new EmailConfigBuilder();
    configure(builder);
    this.receiver.email_configs?.push(builder.build());
    return this;
  }

  build() {
    return this.receiver;
  }
}

export function mockApi(server: SetupServer) {
  return {
    getAlertmanagerConfig: (amName: string, configure: (builder: AlertmanagerConfigBuilder) => void) => {
      const builder = new AlertmanagerConfigBuilder();
      configure(builder);

      server.use(
        rest.get(`api/alertmanager/${amName}/config/api/v1/alerts`, (req, res, ctx) =>
          res(
            ctx.status(200),
            ctx.json<AlertManagerCortexConfig>({
              alertmanager_config: builder.build(),
              template_files: {},
            })
          )
        )
      );
    },

    eval: (response: AlertingQueryResponse) => {
      server.use(
        rest.post('/api/v1/eval', (_, res, ctx) => {
          return res(ctx.status(200), ctx.json(response));
        })
      );
    },
  };
}

export function mockAlertRuleApi(server: SetupServer) {
  return {
    prometheusRuleNamespaces: (dsName: string, response: PromRulesResponse) => {
      server.use(
        rest.get(`api/prometheus/${dsName}/api/v1/rules`, (req, res, ctx) =>
          res(ctx.status(200), ctx.json<PromRulesResponse>(response))
        )
      );
    },
    rulerRules: (dsName: string, response: RulerRulesConfigDTO) => {
      server.use(
        rest.get(`/api/ruler/${dsName}/api/v1/rules`, (req, res, ctx) => res(ctx.status(200), ctx.json(response)))
      );
    },
    rulerRuleGroup: (dsName: string, namespace: string, group: string, response: RulerRuleGroupDTO) => {
      server.use(
        rest.get(`/api/ruler/${dsName}/api/v1/rules/${namespace}/${group}`, (req, res, ctx) =>
          res(ctx.status(200), ctx.json(response))
        )
      );
    },
  };
}

/**
 * Used to mock the response from the /api/v1/status/buildinfo endpoint
 */
export function mockFeatureDiscoveryApi(server: SetupServer) {
  return {
    /**
     *
     * @param dsSettings Use `mockDataSource` to create a faks data source settings
     * @param response Use `buildInfoResponse` to get a pre-defined response for Prometheus and Mimir
     */
    discoverDsFeatures: (dsSettings: DataSourceInstanceSettings, response: PromBuildInfoResponse) => {
      server.use(
        rest.get(`${dsSettings.url}/api/v1/status/buildinfo`, (_, res, ctx) => res(ctx.status(200), ctx.json(response)))
      );
    },
  };
}

// Creates a MSW server and sets up beforeAll, afterAll and beforeEach handlers for it
export function setupMswServer() {
  const server = setupServer();

  beforeAll(() => {
    setBackendSrv(backendSrv);
    server.listen({ onUnhandledRequest: 'error' });
  });

  afterAll(() => {
    server.close();
  });

  return server;
}
