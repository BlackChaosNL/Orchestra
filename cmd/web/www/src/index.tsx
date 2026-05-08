/* @refresh reload */
import './index.css';
import { render } from 'solid-js/web';
import type { Component } from 'solid-js';
import 'solid-devtools';

const App: Component = () => {
  return (
    <p class="text-4xl text-green-700 text-center py-20">Hello tailwind!</p>
  );
};

render(() => <App />, document.getElementById('root')!);
