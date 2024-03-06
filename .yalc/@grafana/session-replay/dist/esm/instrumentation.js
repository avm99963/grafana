import { record } from 'rrweb';
import { BaseInstrumentation, faro, VERSION } from '@grafana/faro-core';
/**
 * Instrumentation for Performance Timeline API
 * @see https://developer.mozilla.org/en-US/docs/Web/API/PerformanceTimeline
 *
 * !!! This instrumentation is in experimental state and it's not meant to be used in production yet. !!!
 * !!! If you want to use it, do it at your own risk. !!!
 */
export class SessionReplayInstrumentation extends BaseInstrumentation {
    constructor() {
        super(...arguments);
        this.name = '@grafana/faro-web-sdk:session-replay';
        this.version = VERSION;
    }
    initialize() {
        this.stopFn = record({
            emit: (event) => {
                var _a, _b;
                faro.api.pushEvent('replay_event', {
                    "timestamp": event.timestamp.toString(),
                    "delay": (_b = (_a = event.delay) === null || _a === void 0 ? void 0 : _a.toString()) !== null && _b !== void 0 ? _b : "",
                    "type": event.type.toString(),
                    "data": JSON.stringify(event.data),
                });
            },
        });
    }
    destroy() {
        this.stopFn();
    }
}
//# sourceMappingURL=instrumentation.js.map