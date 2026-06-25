#include <stdlib.h>

typedef struct IntStack {
  int *items;
  int capacity;
  int size;
} IntStack;

IntStack *stack_new(int capacity) {
  if (capacity <= 0){
    return NULL;
  }
  IntStack *p = malloc(sizeof(IntStack));
  if (p == NULL){
    return NULL;
  }
  p->items = malloc(sizeof(int)*capacity);

  if(p->items == NULL){
    free(p);
    return NULL;
  }
  p->capacity = capacity;
  p->size = 0;
  return p;
  
}

void stack_free(IntStack *stack) {
  if (stack == NULL){
    return;
  }
  free(stack->items);
  free(stack);
  
}

int stack_push(IntStack *stack, int value) {
  if (stack == NULL || stack->size >= stack->capacity){
    return 0;
  }
  stack->items[stack->size] = value;
  stack->size+=1;
    return 1;
  
}

int stack_pop(IntStack *stack, int *out_value) {
  if(stack == NULL || out_value == NULL || stack->size == 0){
    return 0;
  }
  stack->size -=1;
  *out_value = stack->items[stack->size];
  return 1;

  
}

int stack_peek(const IntStack *stack, int *out_value) {
  if(stack == NULL || out_value == NULL || stack->size == 0){
    return 0;
  }
  *out_value = stack->items[stack->size -1];
  return 1;
}

int stack_is_empty(const IntStack *stack) {
  if (stack == NULL){
    return 1;
  }
  if (stack->size == 0){
    return 1;
  }
  return 0;
}
