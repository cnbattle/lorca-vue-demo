<template>
  <div id="app" class="counter-container">
    <h2>{{text}}</h2>
    <h3 class="counter">{{countText}}</h3>
    <div class="btn-row">
      <div class="btn btn-incr" v-on:click="add">+1</div>
      <div class="btn btn-decr" v-on:click="decr">-1</div>
    </div>
  </div>
</template>

<script>
    export default {
        name: 'HelloWorld',
        data() {
            return {
                text: 'Hello lorca and Vue.js',
                countText: "",
            }
        },
        methods: {
            async render() {
                this.countText = `Count: ${await counterValue()}`;
            },
            async add() {
                await counterAdd(1);  // Call Go function, or use `counterAdd(1)`
                this.render()
            },
            async decr() {
                await counterAdd(-1);  // Call Go function, or use `counterAdd(-1)`
                this.render()
            },

        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    user-select: none;
  }

  body {
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #f1c40f;
    font-family: 'Helvetika Neue', Arial, sans-serif;
    font-size: 28px;
  }

  .counter-container {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .counter {
    text-transform: uppercase;
    color: #fff;
    font-weight: bold;
    font-size: 3rem;
  }

  .btn-row {
    display: flex;
    align-items: center;
    margin: 1rem;
  }

  .btn {
    cursor: pointer;
    min-width: 4em;
    padding: 1em;
    border-radius: 5px;
    text-align: center;
    margin: 0 1rem;
    box-shadow: 0 6px #8b5e00;
    color: white;
    background-color: #E4B702;
    position: relative;
    font-weight: bold;
  }

  .btn:hover {
    box-shadow: 0 4px #8b5e00;
    top: 2px;
  }

  .btn:active {
    box-shadow: 0 1px #8b5e00;
    top: 5px;
  }
</style>
