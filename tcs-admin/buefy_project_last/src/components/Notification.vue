<template>
    <section>
        <b-notification
                auto-close type="is-success"
                :active.sync="isActive"
                aria-close-label="Close notification">
            {{message}}
        </b-notification>
        <b-notification
                auto-close type="is-success"
                :active.sync="isActiveEquip"
                aria-close-label="Close notification">
            {{message}}
        </b-notification>
    </section>
</template>

<script>
    export default {
        props: ['message'],
        data() {
            return {
                // isActive: this.$store.state.players.isActive
            }
        },
        computed:{
            isActive(){return this.$store.state.players.isActive},
            isActiveEquip(){return this.$store.state.equips.isActive}
        }
    }
</script>
