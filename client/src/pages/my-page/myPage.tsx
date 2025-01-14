import { useAuth } from "../../provider/auth";
import { TodoList } from "./internal/components/todo-list";
import { TodoCreateForm } from "./internal/components/todo-create-form";
import { useTodos } from "./internal/hook/use-todos";
import { TodoLists } from "./internal/components/todo-lists";
import { useState } from "react";
import { useTodoLists } from "./internal/hook/use-todo-lists";
import { TodoListsCreateForm } from "./internal/components/todo-lists-create-form";

import "./myPage.css";

export const MyPage = () => {
    const { user } = useAuth();
    if (!user) return null;
    const [todoListID,setTodoListID] = useState(user.defaultTodoListID)
    const { todos, todoListName, fetchTodos } = useTodos(todoListID)
    const { todoLists, fetchTodoLists } = useTodoLists()
    
    return (
        <div className="my-page__container">
            <div>
                <h3>{user.name}</h3>
                <TodoListsCreateForm refetch={fetchTodoLists}/>
                <TodoLists todoLists={todoLists} setTodoListID={setTodoListID}/>
            </div>
            <div>
                <h1>{user.name}のマイページ</h1>
                <h2>{todoListName}</h2>
                <TodoCreateForm todoListID={todoListID} refetch={fetchTodos} />
                <TodoList todos={todos} refetch={fetchTodos} />
            </div>
        </div>
    )
}