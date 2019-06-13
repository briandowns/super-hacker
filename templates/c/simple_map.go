package c

// https://github.com/briandowns/simple-map

const SimpleMap = `
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "map.h"

/**
 * contains the value by which the map is 
 * to be increased by when resizing. This 
 */
#define CAPACITY_MULTIPLIER 2

/**
 * list_free frees the memory used for the nodes.
 */
static void
list_free(struct node *n) 
{
    if (!n) {
        return;
    }
    struct node *tmp;
    while (n) {
        tmp = n;
        n = n->next;
        free(tmp);
    }
    free(n);
}

map_t*
map_new(const unsigned int size)
{
    map_t *m = malloc(sizeof(map_t));
    if (!m) {
        return NULL;
    }
    memset(m, 0, sizeof(map_t));
    if (size == 0) {
        m->cap = DEFAULT_SIZE;    
    } else {
        m->cap = size;
    }
    m->len = 0;
    m->list = malloc(sizeof(struct node*)*m->cap);
    memset(m->list, 0, sizeof(struct node*)*m->cap);
    for (int i = 0; i < m->cap; i++) {
        m->list[i] = NULL;
    }
    return m;
}

void
map_free(map_t *m) {
    if (!m) {
        return;
    }
    if (m->list) {
        list_free(*m->list);
    }
    free(m);
}

/**
 * hash hashes the given string to find which
 * bucket it will be placed in.
 */
static int
hash(map_t *m, char *key)
{
    int sum = 0;
    for (int i = 0; i < strlen(key); i++) {
        sum += key[i];
    }
    return sum % m->cap;
}

void*
map_get(map_t *m, char *key)
{
    int pos = hash(m, key);
    struct node *list = m->list[pos];
    struct node *temp = list;
    while (temp) {
        if (strcmp(temp->key, key) == 0) {
            return temp->val;
        }
        temp = temp->next;
    }
    return NULL;
}

/**
 * map_resize will resize the given map with the 
 * new capacity and reinsert the data. 
 */
static int
map_resize(map_t *m, int new_cap) 
{
    map_t *nm = map_new(new_cap);
    if (!nm) {
        return -1;
    }

    for (int i = 0; i < m->len; i++) {
        struct node *list = m->list[i];
        struct node *temp = list;
        while (temp) {
            int st = map_set(nm, temp->key, temp->val);
            if (st != 0) {
                return -1;
            }
            temp = temp->next;
        }
    }

    list_free(*m->list);
    *m = *nm;
    return 0;
}

int
map_set(map_t *m, char *key, void *val)
{
    if (m->len == m->cap) {
        if (map_resize(m, m->cap*CAPACITY_MULTIPLIER) == -1) {
            return -1;
        }
    }
    int pos = hash(m, key);              
    struct node *list = m->list[pos]; 
    struct node *temp = list;                      
    while (temp) {                 
        if (strcmp(temp->key, key) == 0) {
            temp->val = val;
            return 0;
        }
        temp = temp->next;
    }
    struct node *new = malloc(sizeof(struct node)); 
    if (!new) {                                     
        return -1;                                    
    }
    memset(new, 0, sizeof(struct node));
    new->key = strdup(key);
    new->val = val;
    new->next = list;
    m->list[pos] = new;
    m->len++;
    return 0;
}

void
map_del(map_t *m, char *key) {
    int pos = hash(m, key);
    struct node **n = &m->list[pos];
    while (*n) {
        struct node *temp = *n; 
        if (strcmp(temp->key, key) == 0) {
            *n = temp->next;
            break;
        } else {
            temp = (*n)->next;
        }
    }  
    m->len--;
}

int 
map_len(map_t *m) 
{
    int items = 0;
    for (int i = 0; i < m->cap; i++) {
        if (!m->list[i]) {
            items++;
        }
    }
    return items;
}
`
