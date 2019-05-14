<template>
  <div>{{fooCount}}</div>
</template>
<script>
import fooStoreModule from "../store.foo";
export default {
  asyncData({ route, store }) {
    store.registerModule("foo".fooStoreModule);
    return store.dispatch("foo/inc");
  },
  destroyed() {
    this.$store.unregisterModule("foo");
  },
  computed: {
    fooCount() {
      return this.$store.state.foo.count;
    }
  }
};
</script>
