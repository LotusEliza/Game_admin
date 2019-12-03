<template>
    <transition name="fade">
        <section>
        </section>
    </transition>
</template>

<script>
    export default {
        methods: {
            confirm(id) {
                this.$buefy.dialog.confirm({
                    message: 'Are you sure you want to delete equip #'+id+'?',
                    onConfirm: () =>  this.$emit('clicked', id)
                })
            },
        }
    }
</script>

<style>
    /*************Transition******************/
    .fade-enter-active, .fade-leave-active {
        transition: opacity .3s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
        opacity: 0;
    }

</style>
