import { useAuth } from "../../provider/auth";
import { TodoList } from "./internal/components/todo-list";
import { TodoCreateForm } from "./internal/components/todo-create-form";
import { useTodos } from "./internal/hook/use-todos";
import { TodoLists } from "./internal/components/todo-lists";
import { useState } from "react";

export const MyPage = () => {
    const { user } = useAuth();
    if (!user) return null;
    const [todoListID,setTodoListID] = useState(user.defaultTodoListID)
    const { todos, fetchTodos } = useTodos(todoListID)
    
    return (
        <div>
            <div>
                <h3>{user.name}</h3>
                <TodoLists userID={user.id}/>
            </div>
            <div>
                <h1>{user.name}のマイページ</h1>
                <TodoCreateForm todoListID={todoListID} refetch={fetchTodos} />
                <TodoList todos={todos} refetch={fetchTodos} />
            </div>
        </div>
    )
}