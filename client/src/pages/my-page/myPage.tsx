import { useAuth } from "../../provider/auth";
import { TodoList } from "./internal/components/todo-list";
import { TodoCreateForm } from "./internal/components/todo-create-form";
import { useTodos } from "./internal/hook/use-todos";

export const MyPage = () => {
    const { user } = useAuth();
    if (!user) return null;

    console.log(user);
    const { todos, fetchTodos } = useTodos(user.defaultTodoListID)
    
    console.log(todos)
    return (
        <div>
            <h1>{user.name}</h1>
            <TodoCreateForm todoListID={user.defaultTodoListID} refetch={fetchTodos} />
            <TodoList todos={todos} refetch={fetchTodos} />
        </div>
    )
}