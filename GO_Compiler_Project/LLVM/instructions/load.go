package ir

// <result> = load [volatile] <ty>, ptr <pointer>[, align <alignment>][, !nontemporal !<nontemp_node>][, !invariant.load !<empty_node>][, !invariant.group !<empty_node>][, !nonnull !<empty_node>][, !dereferenceable !<deref_bytes_node>][, !dereferenceable_or_null !<deref_bytes_node>][, !align !<align_node>][, !noundef !<empty_node>]
// <result> = load atomic [volatile] <ty>, ptr <pointer> [syncscope("<target-scope>")] <ordering>, align <alignment> [, !invariant.group !<empty_node>]
// !<nontemp_node> = !{ i32 1 }
// !<empty_node> = !{}
// !<deref_bytes_node> = !{ i64 <dereferenceable_bytes> }
// !<align_node> = !{ i64 <value_alignment> }