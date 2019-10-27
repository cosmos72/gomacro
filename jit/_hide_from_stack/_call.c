__attribute__((regparm(0))) void call_trampoline_0b(void *target_ip) {
    typedef void (*func)(void);
    func f = (func)target_ip;
    f();
}
