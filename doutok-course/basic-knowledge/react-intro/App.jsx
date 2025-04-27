import React, { useState, useEffect } from 'react';
import './App.css';

// 一个简单的计数器组件
function Counter() {
  const [count, setCount] = useState(0);
  
  return (
    <div className="counter">
      <h2>计数器: {count}</h2>
      <button onClick={() => setCount(count + 1)}>增加</button>
      <button onClick={() => setCount(count - 1)}>减少</button>
      <button onClick={() => setCount(0)}>重置</button>
    </div>
  );
}

// 一个简单的待办事项组件
function TodoList() {
  const [todos, setTodos] = useState([
    { id: 1, text: '学习 React', completed: false },
    { id: 2, text: '创建一个应用', completed: false },
    { id: 3, text: '部署应用', completed: false }
  ]);
  const [newTodo, setNewTodo] = useState('');
  
  const addTodo = () => {
    if (newTodo.trim() === '') return;
    setTodos([
      ...todos,
      {
        id: Date.now(),
        text: newTodo,
        completed: false
      }
    ]);
    setNewTodo('');
  };
  
  const toggleTodo = (id) => {
    setTodos(
      todos.map(todo => 
        todo.id === id ? { ...todo, completed: !todo.completed } : todo
      )
    );
  };
  
  const deleteTodo = (id) => {
    setTodos(todos.filter(todo => todo.id !== id));
  };
  
  return (
    <div className="todo-list">
      <h2>待办事项列表</h2>
      <div className="add-todo">
        <input 
          type="text" 
          value={newTodo} 
          onChange={(e) => setNewTodo(e.target.value)}
          placeholder="添加新任务"
        />
        <button onClick={addTodo}>添加</button>
      </div>
      
      <ul>
        {todos.map(todo => (
          <li key={todo.id} className={todo.completed ? 'completed' : ''}>
            <span onClick={() => toggleTodo(todo.id)}>{todo.text}</span>
            <button onClick={() => deleteTodo(todo.id)}>删除</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

// 使用 useEffect 的示例组件
function Timer() {
  const [seconds, setSeconds] = useState(0);
  
  useEffect(() => {
    const interval = setInterval(() => {
      setSeconds(seconds => seconds + 1);
    }, 1000);
    
    // 清理函数
    return () => clearInterval(interval);
  }, []);
  
  return <div className="timer">计时器: {seconds} 秒</div>;
}

// 主应用组件
function App() {
  return (
    <div className="app">
      <h1>React 基础示例</h1>
      <Counter />
      <hr />
      <Timer />
      <hr />
      <TodoList />
    </div>
  );
}

export default App; 