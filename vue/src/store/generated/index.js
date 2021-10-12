// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import Amparks100RegistryAmparks100RegistryIdentity from './amparks100/registry/amparks100.registry.identity';
export default {
    Amparks100RegistryAmparks100RegistryIdentity: load(Amparks100RegistryAmparks100RegistryIdentity, 'amparks100.registry.identity'),
};
function load(mod, fullns) {
    return function init(store) {
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: ' + fullns);
        }
        else {
            store.registerModule([fullns], mod);
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns + '/init', null, {
                        root: true
                    });
                }
            });
        }
    };
}
